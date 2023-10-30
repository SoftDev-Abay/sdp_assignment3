package main

type CopyCommand struct {
	textEditor TextEditor
}

func (c *CopyCommand) execute(interface{}) {
	c.textEditor.copy()
}
