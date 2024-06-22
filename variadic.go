package main

import "fmt"

func main() {
	fmt.Println(sumAll(3, 7, 23))

	//jika data berupa slice
	numbers := []int{4, 6, 8}
	fmt.Println(sumAll(numbers...))
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
