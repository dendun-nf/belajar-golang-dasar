package main

import "fmt"

func main() {
	defer endApp()
	runApplication(true)
	fmt.Println("Dundun")
}

func runApplication(error bool) {
	if error {
		panic("Ups something went wrong")
	}

	// tidak akan dieksekusi
	// gunakan recover pada defer function saja
	message := recover()
	fmt.Println("error has been occurred", message)
}

func endApp() {
	fmt.Println("stop application")
	message := recover() // get the return message of panic(msg string)
	fmt.Println("error has been occurred", message)
}
