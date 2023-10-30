package main

type TextEditor interface {
	copy()
	cut()
	paste(int)
	selectText(string)
}
