package main

import (
	"fmt"
	"learn-golang/database"
	_ "learn-golang/internal" //menjalankan blank init
)

func main() {
	fmt.Println(database.GetDatabase())
}
