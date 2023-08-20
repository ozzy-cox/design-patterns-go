package main

import "fmt"

type ITransport interface {
	move()
}

type Vehicle struct{}

func (v *Vehicle) move() {
	fmt.Println(" Vehicle")
}

type OnFoot struct{}

func (o *OnFoot) move() {
	fmt.Println(" onFoot")
}
