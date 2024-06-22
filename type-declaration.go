package main

import "fmt"

func main() {
	type noKTP string

	var userID1 noKTP = "2373475209457043"

	var no_reg string = "124732074274983"
	var userID2 noKTP = noKTP(no_reg)

	fmt.Println("no KTP user 1: ", userID1)
	fmt.Println("no KTP user 2: ", userID2)
}
