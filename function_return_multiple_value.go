package main

import "fmt"

func main() {
	//firstname, lastname := getFullname()
	//fmt.Println(firstname, lastname)

	firstname, _ := getFullname()
	fmt.Println(firstname)
	fmt.Println(getFullname())
}

// di csharp atau lain-lain dikenal dengan tuple
func getFullname() (string, string) {
	return "Deni", "Riyanto"
}
