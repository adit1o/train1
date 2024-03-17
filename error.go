package main

import (
	"errors"
	"fmt"
)

func pembagian(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagian 0 tidak dapat dilakukan")
	}
	return a / b, nil
}
func main() {
	ress, err := pembagian(10, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hasil:", ress)
	}
}
