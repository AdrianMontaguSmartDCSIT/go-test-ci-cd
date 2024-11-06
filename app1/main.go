package main

import (
	"fmt"
)

func main() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(plusOne(1))
}

func addNumbers(a, b int) int {
	return a + b
}

func plusOne(a int) int {
	a = a + 1
	return a
}
