package main

import "fmt"


func checkKelulusan(nilai1 int, nilai2 int) bool {
	var lulusN1 = nilai1 >= 80;
	var lulusN2 = nilai2 >= 80;
	return lulusN1 && lulusN2;
}

func main() {
	var lulus = checkKelulusan(80, 80);
	fmt.Println("Adit=", lulus);
}