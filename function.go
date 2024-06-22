package main

import "fmt"

func main() {
	hello()
	people("erga", "febriawan")
	fmt.Println("hasil operasi", operation(3, 4))

	first, last := getFullname()
	fmt.Println(first, last)

	//ignore return value
	pertama, _ := getFullname()
	fmt.Println(pertama)
}

func hello() {
	fmt.Println("hello")
}

//function parameter
func people(firstName string, lastName string) {
	fmt.Println("you are", firstName, lastName)
}

//function return value
func operation(val1 int, val2 int) int {
	var res int = val1 + val2
	return res
}

//function multiple return value
func getFullname() (string, string) {
	return "erga", "febrian"
}
