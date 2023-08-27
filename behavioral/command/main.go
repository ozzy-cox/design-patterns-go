package main

import "fmt"

type ICommand interface {
	Execute()
}

type PasteCommand struct {
	target IDocument
}

type IDocument interface {
	Paste()
}

type Document struct {
	name string
}

func (d Document) Paste() {
	fmt.Println("Pasted on " + d.name)
}

func (p PasteCommand) Execute() {
	p.target.Paste()
}

type MenuItem struct {
	command ICommand
}

func (m MenuItem) onClick() {
	m.command.Execute()
}

type PasteButton struct {
	MenuItem
}

func main() {
	button := PasteButton{}
	document := Document{"Doc 1"}
	pasteCommand := PasteCommand{}
	pasteCommand.target = document
	button.command = pasteCommand

	button.onClick()
}
