package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"validation error"}
	}

	if id != "erga" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("erga", nil)
	if err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			fmt.Println("Validation error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error:", notFoundErr.Error())
		} else {
			fmt.Println("Unknown Error")
		}
	} else {
		fmt.Println("Success")
	}
}
