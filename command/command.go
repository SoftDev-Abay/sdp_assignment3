package main

type Command interface {
	execute(interface{})
}
