package main

import "fmt"

//alias function type declaration
type Filter func(string) string

func main() {
	sayHelloWithFilter("erga", spamFilter)
	sayHelloWithFilter2("mustafi", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}

func sayHelloWithFilter2(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
