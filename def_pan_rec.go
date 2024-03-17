package main

import (
	"fmt"
)

// defer, panic, recover

// defer    : always di execute di ujung func
// panic    : like thwrow
// recover  : like catch

func logging(s string) {
	fmt.Printf("Logging: %s\n", s)
}

func endApp() {
	fmt.Println("End App")

	errMsg := recover()
	// recover()
	fmt.Println("Error but Continue: ", errMsg)
}

func runApp(isError bool) {
	fmt.Println("runApp...")
	defer endApp()

	if isError == true {
		panic("ups error ~")
	}

	// endApp()

}

func main() {
	runApp(true)
}
