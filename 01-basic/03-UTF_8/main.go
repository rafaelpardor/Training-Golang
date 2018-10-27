package main

import "fmt"

// Printing decimal numbers with his respective string.
func main() {
	fmt.Println("Decimal String")
	fmt.Println("------- \t -----")

	for i := 32; i < 127; i++ {
		fmt.Printf("%d \t %q \n", i, i)
	}
}
