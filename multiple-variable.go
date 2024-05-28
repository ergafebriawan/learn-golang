package main

import "fmt"

func main() {
	var (
		firstName = "erga"
		lastName  = "febriawan"
	)

	const (
		val1 = 3
		val2 = 56
	)

	fmt.Println("nama lengkap: " + firstName + " " + lastName)
	fmt.Println("nilai : ", val1)
}
