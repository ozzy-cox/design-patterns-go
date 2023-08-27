package main

import "fmt"

type IRequestHandler interface {
	HandleRequest()
}

type RequestHandler struct {
	successor IRequestHandler
}

func (r RequestHandler) HandleRequest() {
	if r.successor != nil {
		fmt.Println("forwarded")
		r.successor.HandleRequest()
	} else {
		fmt.Println("Fell off the edge")
	}
}

type Widget struct {
	RequestHandler
}

type Dialog struct {
	Widget
}

type DialogWithHandler struct {
	Widget
}

func (d DialogWithHandler) HandleRequest() {
	fmt.Println("Handled by dialog")
}

type Button struct {
	Widget
}

type Application struct {
	RequestHandler
}

func (a Application) HandleRequest() {
	fmt.Println("Handled by application")
}

func main() {
	dialogButton := Button{}
	dialog := Dialog{}
	application := Application{}

	dialogButton.successor = &dialog
	dialog.successor = &application

	dialogButton.HandleRequest()

	dialogButton = Button{}
	dialogWH := DialogWithHandler{}
	application = Application{}

	dialogButton.successor = &dialogWH
	dialogWH.successor = &application
	dialogButton.HandleRequest()
}
