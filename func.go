package main

import (
	"fmt"
	"strings"
)

// multiple return values
func GetUserData() (string, int) {
	return "Aditya Pratama", 29
}

// named return values
func GetCountry(countryId string) (id int, countryCode string, countryName string) {
	// countryId to lowercase
	countryId = strings.ToUpper(countryId)
	if countryId == "ID" {
		return 0, "id", "Indonesia"
	} else if countryId == "SG" {
		return 1, "sg", "Singapore"
	}
	return 999, "", ""
}

// variadic function
func Sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// function as value
func sayGoodBye(name string) string {
	return "bye " + name
}

// type declaration
type Filter func(w string) string

// function as parameter
func sayWordWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func _blacklistWord(w string) string {
	if w == "Anjing" {
		return "****"
	}
	return w
}

// anonymous function
type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
		return
	}
	fmt.Println("Welcome", name)
}

// recursive function
func factorialLoop(v int) int {
	result := 1
	for i := 0; i < v; i-- {
		result *= (i + 1)
	}
	return result
}

func main() {
	name, age := GetUserData()
	fmt.Println(name, age, "years old")

	// goodbye := sayGoodBye

	sayWordWithFilter("Adit", _blacklistWord)
	filter := _blacklistWord
	sayWordWithFilter("Anjing", filter)

	_, cCode, cName := GetCountry("sg")
	fmt.Printf("Code: %s, Name: %s", cCode, cName)

	fmt.Println(Sum(99, 1))

	fmt.Printf("Goodbye: %s", goodbye("Adit"))

	blackList := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Anjing", blackList)
}
