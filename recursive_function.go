package main

import "fmt"

func main() {
	fmt.Println(10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1)
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}

	return value * factorialRecursive(value-1)
}
