package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("erga")
	if data == nil {
		fmt.Println("data masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}