package main

import "fmt"

func main() {
	// defer function akan selalu dieksekusi
	// setelah sebuah function dieksekusi
	// berhasil atau gagal

	runApp()
}
func logging() {
	fmt.Println("selesai memanggil function")
}
func runApp() {
	// akan tetap dieksekusi walaupun function ini error
	// akan dieksekusi diakhir walaupun di baris pertama
	defer logging()

	fmt.Println("run application")

	// jika error tidak akan dipanggil
	//logging()

}
