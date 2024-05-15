package main

import "fmt"

func main() {
	//var person = map[string]string{}
	//person["name"] = "Deni"
	//person["address"] = "DK Jakarta"
	person := map[string]string{
		"name":    "Deni",
		"address": "DK Jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["tidakada"]) // nilai string default ""
	fmt.Println(person)

	book := make(map[string]string)
	book["tittle"] = "Buku Golang"
	book["author"] = "Deni"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
