package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Deni")
	fmt.Println(result)
	//fmt.Println(helper.sayHello("Deni")) // tidak bisa -> error
}
