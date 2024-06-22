package main

import "fmt"

func main() {
	profile := map[string]string{
		"nama":   "erga",
		"alamat": "biltar",
	}

	fmt.Println(profile["nama"])

	book := make(map[string]string)
	book["title"] = "sapiens"
	book["author"] = "erga"
	book["wrong"] = "ups"

	fmt.Println(book)

	//delete key
	delete(book, "wrong")

	fmt.Println(book)

}
