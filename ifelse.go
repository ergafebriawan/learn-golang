package main

import "fmt"

func main() {
	name := "ega"

	if name == "erga" {
		fmt.Println("nama yang benar")
	} else if name == "febri" {
		fmt.Println("hampir benar")
	} else {
		fmt.Println("nama yang salah")
	}

	//if short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah sesuai")
	}
}
