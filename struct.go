package main

import "fmt"

//selalu diawali dengan huruf besar
type Profile struct {
	Name    string
	Address string
	Gender  string
	Age     int
}

func (profile Profile) sayHello() {
	fmt.Println("Hello", profile.Name)
}

func main() {
	var user1 Profile
	user1.Name = "erga febriawan"
	user1.Address = "biltar"
	user1.Gender = "Man"
	user1.Age = 26

	user2 := Profile{
		Name:    "mustafi",
		Address: "Blitar",
		Gender:  "man",
		Age:     32,
	}
	fmt.Println(user1.Name)
	fmt.Println(user2)

	user2.sayHello()
}
