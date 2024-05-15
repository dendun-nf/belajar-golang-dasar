package main

import "fmt"

func main() {
	deni := Man{"Deni"}

	deni.Married()

	fmt.Println(deni)
}

type Man struct {
	Name string
}

// tanpa asterisk operator maka sebenarnya man hanyalah copy dari structnya
// maka perubahan pada parameter man type Man tidak akan mengubah apapun
// ketika dipanggil oleh object dari struct itu sendiri
// dengan man *Man maka data pada 'man' akan mengacu pada struct yang memanggil function ini sendiri
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
