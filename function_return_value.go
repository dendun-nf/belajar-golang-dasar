package main

import "fmt"

func main() {
	result := getHello("Deni")
	fmt.Println(result)
	fmt.Println(getHello("Budi"))
}

func getHello(name string) string {
	return "Hello " + name
}
