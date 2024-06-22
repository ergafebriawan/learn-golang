package main

import "fmt"

func main() {
	a, b, c := getCompleteName()
	fmt.Println("nama anda:", a, b, c)
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "erga"
	middleName = "febrian"
	lastName = "jaya"

	return firstName, middleName, lastName
}
