package helper

import "fmt"

// / Access Modifier
// `PrintHello` bisa diakses di packages lain
// `printHello` tidak bisa diakses dipackages lain
func PrintHello(name string) {
	fmt.Println("Hello", name)
}