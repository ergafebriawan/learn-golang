package main

import "fmt"

func main() {
	var user = [...]string{
		"gise",
		"erga",
		"nisa",
		"bintang",
		"senja",
	}

	slice1 := user[2:5]
	fmt.Println(slice1)
	fmt.Println(len(slice1))

	slice2 := user[:3]
	fmt.Println(slice2)

	slice3 := user[3:]
	fmt.Println(slice3)
	fmt.Println(cap(slice3))

	slice4 := user[:]
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]

	fmt.Println(daySlice1)
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	//append slice
	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println(daySlice2)

	//new slice
	var newslice []string = make([]string, 3, 6)
	newslice[0] = "king"
	newslice[1] = "kong"
	newslice[2] = "gorrila"

	fmt.Println(newslice)

	//copy slice

	fromSlice := days[3:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	//perbedaan array dan slice
	iniArray := [...]int{1, 3, 2, 7, 6}
	iniSlice := []int{3, 5, 6, 9, 1}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
