package main

import "fmt"

func main() {
	name := "Deni Riyanto"
	fmt.Println(name)

	name = "Deni Dendun"
	fmt.Println(name)

	// jika tidak digunakan akan error
	var (
		firstname = "Deni"
		lastname  = "Riyanto"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
}
