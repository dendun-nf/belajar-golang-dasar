package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error:", err)
	}
}
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan NOL")
	}
	return nilai / pembagi, nil
}
