package main

import "fmt"

func main() {
	sayHelloTo("Deni", "Riyanto")
	sayHelloTo("Budi", "Nugraha")
}

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}
