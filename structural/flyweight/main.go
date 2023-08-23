package main

import (
	"fmt"
	"strings"
)

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func (ff *FlyweightFactory) getFlyweight(key string) Flyweight {
	flyweight, exists := ff.flyweights[key]

	if exists {
		return flyweight
	}

	newFlyweight := Flyweight{code: key}

	ff.flyweights[key] = newFlyweight

	return newFlyweight
}

type Flyweight struct {
	code string
}

// Pass in extrinsic state
func (f *Flyweight) Draw(isCapitalized bool) {
	if isCapitalized {
		fmt.Println(strings.ToUpper(f.code))
	} else {
		fmt.Println(f.code)
	}
}

func main() {
	factory := FlyweightFactory{}
	factory.flyweights = make(map[string]Flyweight)
	a := factory.getFlyweight("a")
	b := factory.getFlyweight("b")
	c := factory.getFlyweight("c")
	d := factory.getFlyweight("d")

	a.Draw(false)
	b.Draw(true)
	c.Draw(false)
	d.Draw(true)
}
