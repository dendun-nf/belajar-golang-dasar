package main

import "fmt"

func main() {
	//const firstname string = "Deni"
	//const lastname = "Riyanto"

	const (
		firstname = "Deni"
		lastname  = "Riyanto"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

	// akan error
	//firstname = "Dundun"
	//lastname = "Dondon"

}
