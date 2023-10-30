package main

type CutCommand struct {
	textEditor TextEditor
}

func (c *CutCommand) execute(interface{}) {
	c.textEditor.cut()
}
