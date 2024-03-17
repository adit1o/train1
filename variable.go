package main

import "fmt"

var (
	isCorrect 	bool
	nickName	string
)


func main() {
	
	isCorrect = false;
	nickName = "adit"
	namaku := "Aditya Pratama";



	fmt.Println("isCorrect", isCorrect);
	fmt.Println("name", nickName);
	fmt.Println("namaku:", namaku);


	// change
	nickName = "prtm"
	fmt.Println("name", nickName);


	fmt.Println(nickName[0])

	var asd = string(namaku[0]);
	fmt.Println(asd)

}
