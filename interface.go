package main

import (
	"fmt"
	"main/helper"
)

/// abstraction ---- contract ---- prototype

type HasName interface {
	GetName() string
	// GetStatus() bool
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func (p Animal) GetName() string {
	return p.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func OpoIki() any {
	return "1"
}

func main() {

	adit := Person{"Adit"}

	kucing := Animal{"Kucing"}

	SayHello(adit)
	SayHello(kucing)

    fmt.Println(OpoIki())

}
