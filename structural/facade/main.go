package main

import "fmt"

// Facade
type Compiler struct{}

// This kind of seems like a strategy pattern almost
func (co *Compiler) compile(s Scanner, p Parser, n ProgramNodeBuilder, c CodeGenerator) {
	// Complex logic
	s.Scan()
	p.Parse()
	n.BuildNode()
	c.GenerateCode()
}

type (
	Scanner            struct{}
	Parser             struct{}
	ProgramNodeBuilder struct{}
	CodeGenerator      struct{}
)

func (s *Scanner) Scan() {
	fmt.Println("Scan")
}

func (s *Parser) Parse() {
	fmt.Println("Parse")
}

func (s *ProgramNodeBuilder) BuildNode() {
	fmt.Println("BuildNode")
}

func (s *CodeGenerator) GenerateCode() {
	fmt.Println("GenerateCode")
}

func main() {
	compiler := Compiler{}

	compiler.compile(Scanner{}, Parser{}, ProgramNodeBuilder{}, CodeGenerator{})
}
