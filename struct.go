package main

import "fmt"

func main() {
	var deni Customer
	deni.Name = "Deni"
	deni.Address = "DK Jakarta"
	deni.Age = 20

	fmt.Println(deni)
	fmt.Println(deni.Name)
	fmt.Println(deni.Address)
	fmt.Println(deni.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "DI Yogyakarta",
		Age:     21,
	}

	budi := Customer{"Budi", "Aceh", 22}
	fmt.Println(joko)
	fmt.Println(budi)

	deni.sayHello("Budi")
	budi.sayHello("Agus")
	joko.sayHello("Deni")
}

// kumpulan field dengan primitive / function type
type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello,", customer.Name, "My names is", name)
}
