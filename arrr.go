package main

import "fmt"

// / better use slice than array coz slice dynamis
func main() {
	// late
	var names [2]string
	names[0] = "Aditya"
	names[1] = "Pratama"

	var zoo = []string{"kucing","tikus"}

	zoo = append(zoo, "kambing")
	fmt.Println("zoo", zoo)



	
	
	
	fmt.Println(names[0])
	fmt.Println(names[1])

	// init
	var names2 = [2]string{"Aditya", "Pratama"}
	fmt.Println(names2)

	// func array
	var cities = [...]string{"aceh", "medan", "jakarta", "bandung"}
	fmt.Println(cities)

	// slicing
	var month = [...]string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	var slice1 = month[4:7]
	var slice2 = month[:4]
	var slice3 = month[4:]
	fmt.Println(slice1, "length=", len(slice1))
	fmt.Println(slice2, "length=", len(slice2))
	fmt.Println(slice3, "length=", len(slice3))

	// append
	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	days = append(days, "newDay")
	days[len(days)-1] = "edit"
	fmt.Println(days, len(days))

	// make
	var zodiac = make([]string, 12)
	zodiac[0] = "capricorn"
	zodiac[1] = "aquarius"
	zodiac[2] = "pisces"
	zodiac[3] = "aries"
	zodiac[4] = "taurus"
	zodiac[5] = "gemini"
	zodiac[6] = "cancer"
	zodiac[7] = "leo"
	zodiac[8] = "virgo"
	zodiac[9] = "libra"
	zodiac[10] = "scorpio"
	zodiac[11] = "sagittarius"
	fmt.Println(zodiac[10:])

	// copy
	var newZodiac = make([]string, len(zodiac))
	copy(newZodiac, zodiac)
	newZodiac[11] = "capricorn"
	fmt.Println(newZodiac)
}
