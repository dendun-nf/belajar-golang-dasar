package main

import "fmt"

func main() {
	var alamat1 = new(Address) // sama dengan &Address{}
	var alamat2 = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

type Address struct {
	City, Province, Country string
}
