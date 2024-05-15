package main

import "fmt"

func main() {
	//address1 := Address{"Jakarta Pusat", "DK Jakarta", "Indonesia"}
	////address2 := Address{"Johor Baru", "Kuala Lumpur", "Malaysia"}
	//
	//// pass by value, address1 copied, stored to address2
	//address2 := address1
	//
	//fmt.Println(address1)
	//fmt.Println(address2)
	//
	//// any changes on address2 didn't affect address1
	//address2.Province = "Aceh"
	//fmt.Println(address2) // aceh
	//fmt.Println(address1) // still DK Jakarta

	var address1 = Address{"Jakarta Pusat", "DK Jakarta", "Indonesia"}
	var address2 = &address1

	address2.Province = "Aceh"

	fmt.Println(address1)
	fmt.Println(address2)

}

func changeVal(val int) {
	val = 0
}

type Address struct {
	City, Province, Country string
}
