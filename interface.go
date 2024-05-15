package main

import "fmt"

func main() {
	deni := Person{"Deni"}
	SayHello(deni)
	animal := Animal{"Chuu"}
	SayHello(animal)
}

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}
