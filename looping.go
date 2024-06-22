package main

import "fmt"

func main() {
	//example 1
	counter := 1

	for counter <= 10 {
		fmt.Println("value: ", counter)
		counter++
	}

	//example 2
	for count := 1; count <= 10; count++ {
		fmt.Println("nilai: ", counter)
	}

	//for range
	users := []string{"erga", "senja", "nidavellir"}

	for index, user := range users {
		fmt.Println("index ke ", index, " value ", user)
	}
}
