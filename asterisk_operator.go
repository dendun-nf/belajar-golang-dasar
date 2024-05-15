package main

import "fmt"

func main() {

	var address1 = Address{"Jakarta Pusat", "DK Jakarta", "Indonesia"}
	var address2 = &address1

	address2.Province = "Aceh"
	fmt.Println(address1)
	fmt.Println(address2)

	// mendefinisikan address2 sebagai pointer baru dari sebuah Address
	//address2 = &Address{"Jakarta Pusat", "DK Jakarta", "Indonesia"} // mendefiniskan ulang

	/**
	mengakses data dari pointer yang ada pada address2
	dengan * untuk didefinisikan ulang sebagai Address{}
	yang mengakibatkan data di address1 ikut berubah
	ex: anggap ada sebuah address
	address1 = 0x00 (Address struct)
	address2 = &address1 = 0x00

	tanpa menggunakan operator *
	&Address = 0x04
	address2 = &Address{bla, bla, bla} // pointer lain dengan value dari &Address yang baru
	jika menggunakan operator *

	address1 = 0x00
	*address2 = 0x00
	jika didefinisikan dengan:
	*address2 = Address{bla,bla,bla}
	maka ini sama dengan:
	0x00 = Address{bla,bla,bla}
	yang dimana:
	0x00 = address1 = address2
	maka address1 ikut berubah dengan pendifinisian ulang dari asterisk operator *address2 = Address{}

	sehingga
	*/
	*address2 = Address{"Jakarta Pusat", "DK Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}

type Address struct {
	City, Province, Country string
}
