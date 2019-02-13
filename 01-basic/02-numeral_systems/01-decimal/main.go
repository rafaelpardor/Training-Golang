package main

import "fmt"

func main() {
	fmt.Println("Integers numbers")
	fmt.Println("-------")

	for i := 1; i < 11; i++ {
		fmt.Printf("%d\n", i) // We use %d to print int numbers.
	}
}
