package main

import "fmt"

// * == asterisk
// & == reference
type Users struct {
	Name string
	Age  int
}

func changeUserName(user *Users, newName string) {
	user.Name = newName
}

func main() {

	user1 := Users{"Adit", 20}
	// user2 := user1 // copy only (not reference to user1)
	user2 := &user1 // reference to user1 (pointer) ==== `User*`
	user2.Name = "Adiittt"

	/// arterisc operation
	// user2 = &Users{"Aditya Pratama", 29}
	// *user2 = Users{"Aditya Pratama", 29}

	/// new keyword operator
	user2 = new(Users)

	changeUserName(user2, "Adit Psgt")
	fmt.Println(user1, user2)
}
