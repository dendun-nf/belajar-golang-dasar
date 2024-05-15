package main

import "fmt"

func main() {
	result := random()

	//resultStr := result.(string)
	//fmt.Println(resultStr)
	//resultInt := result.(int) // error
	//fmt.Println(resultInt)    // tidak akan dieksekusi

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}

func random() any {
	return "Ok"
}
