package main

import "fmt"

// An array is a numbered sequence of elements of a single type with a fixed length.
func main() {
	var x [10]int // We create an array like this...

	fmt.Println(x)                              // We print all the array.
	fmt.Println("Lenght of the array:", len(x)) // We print the lenght of the array using the function 'len()'
	fmt.Printf("Type-%T\n", x)                  // We print the type using %T

	x[9] = 777
	fmt.Println(x[9]) // We print an especific element of the array <array>[n]
}
