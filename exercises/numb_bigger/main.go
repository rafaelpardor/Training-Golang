package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program will print the biggest number among 3 numbers")
	sortNumbers(10, 20, 3)
}

func sortNumbers(A, B, C int) (int, int, int) {
	if A > B {
		if A > C {
			fmt.Printf("%d is bigger than %d and %d.\n", A, B, C)
		} else {
			fmt.Printf("%d is bigger than %d and %d.\n", C, A, B)
		}
	} else {
		if B > C {
			fmt.Printf("%d is bigger than %d and %d.\n", B, B, C)
		} else {
			fmt.Printf("%d is bigger than %d and %d.\n", C, A, B)
		}
	}
	return A, B, C
}
