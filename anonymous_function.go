package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		if name == "anjing" {
			return true
		}
		return false
	}
	registerUser("Deni", blacklist)
	registerUser("anjing", blacklist)
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
