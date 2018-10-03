package main

import (
	"fmt"
)

// We make a for to show the type and the value of float numbers.
func main() {
	fmt.Println("By default, a float variable is a float64.")

	for i := 1.10; i < 11.12; i++ {
		fmt.Printf("The variable we have been created, is a %T with a value of %#v.\n", i, i)
	}
}
