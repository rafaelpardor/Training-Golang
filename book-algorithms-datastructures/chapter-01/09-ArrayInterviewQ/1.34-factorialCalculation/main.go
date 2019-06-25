package main

import "fmt"

// Factorial Calculation N!=N*(N-1)
func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}

func main() {
	n := 2
	fmt.Printf("Factorial of %d is: %d\n", n, Factorial(n))
}
