package main

import (
	"fmt"
)

// We make a for to show the type and the value of int numbers.
func main() {
	for i := 1; i < 11; i++ {
		fmt.Printf("The variable we have been created is a %T, with a value of %v.\n", i, i)
	}
}
