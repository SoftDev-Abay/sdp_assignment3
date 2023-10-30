package main

import "fmt"

func main() {
	mousePad := &MousePad{text: "Go is a statically typed, compiled programming language designed at Google by Robert Griesemer."}
	selectCommand := &SelectCommand{
		textEditor: mousePad,
	}

	copyCommand := &CopyCommand{
		textEditor: mousePad,
	}

	cutCommand := &CutCommand{
		textEditor: mousePad,
	}

	pasteCommand := &PasteCommand{
		textEditor: mousePad,
	}

	selectHotKey := &HotKey{
		command: selectCommand,
	}

	copyHotKey := &HotKey{
		command: copyCommand,
	}

	cutHotKey := &HotKey{
		command: cutCommand,
	}

	pasteHotKey := &HotKey{
		command: pasteCommand,
	}

	fmt.Println("MousePad text:", mousePad.text)
	selectHotKey.press("Robert Griesemer")

	copyHotKey.press(nil)

	cutHotKey.press(nil)

	fmt.Println("MousePad text:", mousePad.text)

	selectHotKey.press("Abay Aliyev")

	copyHotKey.press(nil)

	pasteHotKey.press(len(mousePad.text) - 1)

	fmt.Println("MousePad text:", mousePad.text)

}
