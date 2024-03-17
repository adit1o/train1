package main

import "fmt"



func main() {

	/// basic 
	var a, b, c int
	a = 10
	b = 20
	c = a + b
	fmt.Println("a + b =", c)


	/// augmented assignment
	k := 10
	k += 20
	fmt.Println("k =", k)

	/// unary operator
	n := 10
	n++
	fmt.Println("n =", n);

	bo := false;
	bo = !bo
	fmt.Println("bo =", bo)

}