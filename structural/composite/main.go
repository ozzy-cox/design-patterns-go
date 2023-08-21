package main

import "fmt"

type IComponent interface {
	Draw()
}

type Node struct {
	name string
}

func (n *Node) Draw() {
	fmt.Println(n.name + " is drawn")
}

type Composite struct {
	nodes []Node
}

func (c *Composite) Draw() {
	for _, node := range c.nodes {
		node.Draw()
	}
}

func main() {
	leaf1 := Node{"leaf1"}
	leaf2 := Node{"leaf2"}
	leaf3 := Node{"leaf3"}

	comp := Composite{}
	comp.nodes = make([]Node, 0)

	comp.nodes = append(comp.nodes, leaf1)
	comp.nodes = append(comp.nodes, leaf2)
	comp.nodes = append(comp.nodes, leaf3)
	leaf1.Draw()
	comp.Draw()
}
