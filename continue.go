package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // langsung ke iterasi berikutnya tanpa eksekusi code setelah ini
		}
		fmt.Println("Perulangan ke ", i)
	}
}
