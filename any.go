package main

import "fmt"

func main() {
	kosong := Ups()

	fmt.Println(kosong)
}

//func Ups() interface{} {
//	return 1
//}

func Ups() any {
	return 1
	//return true
	//return ""
}
