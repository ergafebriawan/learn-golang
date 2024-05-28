package main

import "fmt"

func main() {
	var nilai int16 = 34
	var nilai_conv int32 = int32(nilai)
	var name string = "erga"
	var str_conv string = string(name[1])

	fmt.Println("hasil konversi: ", nilai_conv)
	fmt.Println("hasil konversi string: ", str_conv)
}
