package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//pass by value
	address1 := Address{"Bltar", "Jawa Timur", "Indonesia"}
	address2 := address1

	address2.City = "Malang"
	fmt.Println(address1)
	fmt.Println(address2)

	// pointer
	var address12 Address = Address{"Bltar", "Jawa Timur", "Indonesia"}
	var address22 *Address = &address1

	address2.City = "Malang"
	fmt.Println(address12)
	fmt.Println(address22)

}
