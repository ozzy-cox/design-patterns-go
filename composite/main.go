package main

import (
	"fmt"
)

type IComponent interface {
	Draw()
	Add(c *IComponent)
	Remove() (*IComponent, error)
	getChild(i int) (IComponent, error)
}

type Graphic struct {
	name string
}

func (c Graphic) Draw() {
	fmt.Println(c.name)
}

type (
	Rectangle struct {
		component Graphic
	}
	Line struct {
		component Graphic
	}
	Text struct {
		component Graphic
	}
)

func (c Line) Draw() {
	fmt.Println("Line")
}

func (c Text) Draw() {
	fmt.Println("Text")
}

func (c Rectangle) Draw() {
	fmt.Println("Rectangle")
}

type Picture struct {
	component Graphic
	children  []IComponent
}

func (c Picture) Draw() {
	fmt.Println("picturestart")
	for _, child := range c.children {
		child.Draw()
	}
	fmt.Println("pictureend")
}

func TestComposite() {
	

}
