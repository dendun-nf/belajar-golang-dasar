package main

import "fmt"

func main() {
	name := "Riyanto"

	if name == "Deni" {
		fmt.Println("Hello Deni")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")

	} else if name == "Deni" {
		fmt.Println("Lah Deni Lagi") // tidak akan dieksekusi
	} else {
		fmt.Println("Hallo, boleh kenalan?")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
