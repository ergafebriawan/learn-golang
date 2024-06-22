package main

import "fmt"

func main() {
	name := "erga"

	switch name {
	case "erga":
		fmt.Println("halo erga")
	case "febri":
		fmt.Println("halo febri")
	default:
		fmt.Println("invalid name")
	}

}
