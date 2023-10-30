package main

type SelectCommand struct {
	textEditor TextEditor
}

func (s *SelectCommand) execute(text interface{}) {
	s.textEditor.selectText(text.(string))
}
