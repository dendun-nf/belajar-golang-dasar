package main

import "fmt"

func main() {
	name := "Deni"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Deni":
		fmt.Println("Hello Deni")
	default:
		fmt.Println("Hello, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")

	}

	name = "Riyanto"
	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 10:
		fmt.Println("Nama kepanjangan")
	default:
		fmt.Println("Nama sudah benar")

	}
}
