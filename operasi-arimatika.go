package main

import "fmt"

func main() {
	var a = 3
	var b = 5
	var c = 2

	var res = a + b - c

	//augmented assigment
	res += 10

	//unary operator
	res++

	fmt.Println("hasil: ", res)
}
