package main

import "fmt"

func main() {
	hello := getHello

	fmt.Println(hello("erga"))
}

func getHello(name string) string {
	return "hello " + name
}
