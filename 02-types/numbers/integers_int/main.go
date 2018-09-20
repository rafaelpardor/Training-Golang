package main

import (
	"fmt"
)

func main() {
	fmt.Println("-.- Integers -.-")
	ints()
	fmt.Println(getIntegers(1, 2, 3))
	fmt.Println(sum(1, 2))
}

func ints() {
	fmt.Print("\n")

	var x int32 = 71

	fmt.Printf("Thr variable we have been created, is a%T with a value of %v.\n", x, x)
}

func getIntegers(a, b, c int) (int, int, int) {
	fmt.Print("\n")

	return a, b, c
}

func sum(a, b int) int {
	fmt.Print("\n")

	return a + b
}
