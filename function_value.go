package main

import "fmt"

func main() {
	contoh := getGoodBye
	misal := getGoodBye
	fmt.Println(contoh("Dundun"))
	fmt.Println(misal("Agus"))
}

func getGoodBye(name string) string {
	return "Goodbye " + name
}
