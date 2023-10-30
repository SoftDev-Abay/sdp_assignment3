package main

type PasteCommand struct {
	textEditor TextEditor
}

func (p *PasteCommand) execute(position interface{}) {
	p.textEditor.paste(position.(int))
}
