package main

import (
	"fmt"
	"learn-golang/helper"
)

func main() {
	fmt.Println("open")
	result := helper.SayHello("erga")
	fmt.Println(result)
}
