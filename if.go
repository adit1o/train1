package main

import "fmt"

func main() {
	country := "Indonesia"

	if country == "Indonesia" {
		fmt.Println("Halo Jakarta")
	} else {
		fmt.Println("Halo dunia")
	}

	
	if l := len(country); l == 9 {
		fmt.Println("Halo Indonesia")
	}

}
