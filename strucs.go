package main

import "fmt"

type UserProfile struct {
	Name, Address string
	Age           int
	isJomblo      bool
}

// / struct method
func (userP UserProfile) sayHello(yourName string) {
	fmt.Println("Hello "+yourName+", my name is", userP.Name)

}

func main() {

	adit := UserProfile{
		Name:     "Adit",
		Address:  "Jl. Raya Cibadak",
		Age:      20,
		isJomblo: true,
	}

	budi := UserProfile{
		Name:     "Budi",
		Address:  "Jl. Indonesia Raya",
		Age:      39,
		isJomblo: false,
	}

	coki := UserProfile{"Coki", "Jl. Anak Desa", 19, true}

	fmt.Println("data user adit: ", adit)
	fmt.Println("data user budi: ", budi)
	fmt.Println("data user budi: ", coki)

	adit.sayHello("joko")

}
