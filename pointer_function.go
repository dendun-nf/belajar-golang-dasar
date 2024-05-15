package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address{}
	changeCountryToIndonesia(address)
	fmt.Println(address) // tetap kosong, tidak berubah

	changeCountryToIndonesiaByReference(&address)
	fmt.Println(address)
}

func changeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func changeCountryToIndonesiaByReference(address *Address) {
	address.Country = "Indonesia"
}
