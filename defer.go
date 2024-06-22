package main

import (
	"fmt"
)

// defer
func logging() {
	fmt.Println("selesai memangil function")
}

func runApplication() {
	defer logging()
	fmt.Println("run application")
}

// panic
func endApp() {
	fmt.Println("Application End")
	//recover
	message := recover()
	fmt.Println("terjadi panic:", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	runApp(true)
	fmt.Println("erga febriawan")
}
