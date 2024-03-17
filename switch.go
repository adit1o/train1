package main

import "fmt"

func main() {

    age := 28


    switch age {
    case 17:
        fmt.Println("belum cukup umur")
    case 21:
        fmt.Println("OKEEEEE")
    default:
        fmt.Println("kamu sudah cukup umur")
    }
}