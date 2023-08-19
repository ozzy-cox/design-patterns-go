package main

import "fmt"

type DBConnection struct{}

var instance *DBConnection

func getInstance() *DBConnection {
	// Consider locking for multithreaded usage
	// Refer to here : https://refactoring.guru/design-patterns/singleton/go/example
	if instance == nil {
		instance = &DBConnection{}
		fmt.Println("instantiated")
	} else {
		fmt.Println("not instantiated")
	}

	return instance
}

func main() {
	inst := getInstance()
	fmt.Println(inst)
	inst2 := getInstance()
	fmt.Println(inst2)
}
