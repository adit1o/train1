package main

import "fmt"

type Personn struct {
	Name string
}

func (g *Personn) Men() {
	g.Name = "Mr. " + g.Name
}

func (g *Personn) Women() {
	g.Name = "Miss. " + g.Name
}

func main() {
	gender := Personn{"Adit"}
	gender.Women()
	fmt.Println(gender)
}
