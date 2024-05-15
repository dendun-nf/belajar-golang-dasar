package main

import "fmt"

func getCompleteName() (firstname, middlename, lastname string) {

	firstname = "Deni"
	middlename = "Dundun"
	lastname = "Riyanto"

	return firstname, middlename, lastname
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
