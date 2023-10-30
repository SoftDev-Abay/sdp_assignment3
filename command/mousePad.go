package main

import (
	"fmt"
	"strings"
)

type MousePad struct {
	text         string
	selectedText string
	savedText    string
}

func (m *MousePad) selectText(text string) {
	m.selectedText = text
	fmt.Printf("Selected text - %v \n", m.selectedText)
}

func (m *MousePad) copy() {
	if len(m.selectedText) <= 0 {
		return
	}
	m.savedText = m.selectedText
	fmt.Printf("Text copied - %v \n", m.savedText)
}

func (m *MousePad) cut() {
	if len(m.selectedText) <= 0 {
		return
	}
	textToFind := m.selectedText
	text := m.text
	posOfTextToCut := strings.Index(text, textToFind)
	if posOfTextToCut == -1 {
		return
	}

	first := text[:posOfTextToCut]
	second := text[posOfTextToCut+len(textToFind):]
	newText := first + second
	m.text = newText

	fmt.Printf("Text cut - %v \n", m.selectedText)
}

func (m *MousePad) paste(pos int) {
	if len(m.savedText) <= 0 || pos < 0 {
		return
	}
	text := m.text
	first := text[:pos] + m.savedText
	second := text[pos:]
	newText := first + second
	m.text = newText

	fmt.Printf("Text pasted - %v \n", m.savedText)
}
