package main

import "fmt"

func main() {

	/// like another language
	months := []string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	for i := 0; i < len(months); i++ {
		fmt.Println(months[i])
	}

	/// newbie
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	/// basic
	counter := 1
	for counter <= 10 {
		fmt.Println("counter = ", counter)
		counter++
	}

	/// range
	for i, month := range months {
		fmt.Println((i + 1), month)
	}

    for _, month := range months {
        if month == "april" {
            continue
        }
        fmt.Println(month)
    }
}
