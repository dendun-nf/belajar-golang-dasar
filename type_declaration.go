package main

import "fmt"

func main() {
	type NoKTP string

	var ktpDeni NoKTP = "11111111"

	var contoh = "222222"
	var contohKtp = NoKTP(contoh)

	fmt.Println(contohKtp)
	fmt.Println(ktpDeni)
}
