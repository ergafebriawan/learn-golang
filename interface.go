package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(has HasName) {
	fmt.Println("Hello", has.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{Name: "erga"}
	sayHello(person)
}
