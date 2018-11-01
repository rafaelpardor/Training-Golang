package main

import "fmt"

// Printing decimal numbers with his respective string.
func main() {
	fmt.Println("Decimal String")
	fmt.Println("------- ------")

	for i := 32; i < 127; i++ {
		fmt.Printf("%d \t %q \n", i, i) // Using %d to print int, and %q to a single-quoted character literal
	}
}
