package main

type HotKey struct {
	command Command
}

func (k *HotKey) press(param interface{}) {
	k.command.execute(param)
}
