package main

import "fmt"

func rand() any {
	return false
}

func main() {

	var rr = rand()
	// if b, ok := rr.(int); ok {
	//     fmt.Println(b)
	// }

	switch rr.(type) {
	case string:
		fmt.Println("jadi string", rr)
	case int:
		fmt.Println("jadi int", rr)
	case bool:
		fmt.Println("jadi bool", rr)
	default:
		fmt.Println("unknow", rr)
	}

}
