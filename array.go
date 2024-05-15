package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Deni"
	names[1] = "Riyanto"
	names[2] = "Dendun"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{90, 80, 95, 100, 110}
	//var values = [...]int{}
	fmt.Println(values)
	fmt.Println(values[2])

}
