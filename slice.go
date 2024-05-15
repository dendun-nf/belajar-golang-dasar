package main

import "fmt"

func main() {
	names := [...]string{"Deni", "Riyanto", "Joko", "Budi", "Nugraha", "Joe"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:] // sabtu. minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Deni"
	newSlice[1] = "Deni"
	//newSlice[2] = "Dendun" // akan error

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Deni")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Dundun"
	fmt.Println(newSlice2)
	fmt.Println(newSlice) // newSlice terpengaruh karena modifikasi dari newSlice2[0] = "Dundun"

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// hati hati saat membuat slice pada array

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
