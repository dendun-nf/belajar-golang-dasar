package main

import "fmt"

func main() {
	sayHelloWithFilter("Deni", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Deni", filter)
}

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

type Filter func(string) string
