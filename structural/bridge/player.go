package main

import "fmt"

// Abstraction
type ICharacter interface {
	travel()
}

type Character struct {
	name      string
	transport ITransport
}

func (c *Character) travel() {
	fmt.Print(c.name + " moving with")
	c.transport.move()
}

// Refined abstraction
type Player struct {
	Character
}

type NPC struct {
	Character
}
