package main

import (
	"fmt"
	"strings"
)

type ValiadationError struct {
	message string
}

type NotFoundError struct {
	message string
}

func (v *ValiadationError) Error() string {
	return v.message
}

func (v *NotFoundError) Error() string {
	return v.message
}

func login(username string) error {
	username = strings.ToLower(username)

	if username == "" {
		return &ValiadationError{"username cannot be empty"}
	}

	if username != "adit" {
		return &NotFoundError{"username not found"}
	}

	return nil
}

func main() {
	login := login("Budi")
	if login != nil {
		switch login.(type) {
		case *ValiadationError:
			fmt.Println(login, "--> ValiadationError")
		case *NotFoundError:
			fmt.Println(login, "--> NotFoundError")
		}
	} else {
		
		fmt.Println("login success")
	}
}
