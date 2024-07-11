package mypackage

import "fmt"

func Sum(a, b int) int {
	fmt.Printf("Sum %d + %d \n", a, b)
	return a + b
}

func Times(a, b int) int {
	fmt.Printf("Times %d * %d \n", a, b)
	return a * b
}
