package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "mustafi"
	name[1] = "senja"
	name[2] = "perompak"

	var val = [3]int{
		43,
		56,
		22,
	}

	var data = [...]int{
		43,
		56,
		22,
		75,
	}

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	fmt.Println(val[0])

	fmt.Println(len(data))
}
