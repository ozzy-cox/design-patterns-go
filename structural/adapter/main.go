package main

import "fmt"

type IAdapter interface {
	specificRequest()
}

type Adapter struct {
	adaptee Adaptee
}

type Adaptee struct{}

func (a Adaptee) request() {
	fmt.Println("request")
}

func (a Adapter) specificRequest() {
	fmt.Println("spec")
	a.adaptee.request()
}

func main() {
	adapter := Adapter{}
	adapter.adaptee = Adaptee{}

	adapter.specificRequest()
}
