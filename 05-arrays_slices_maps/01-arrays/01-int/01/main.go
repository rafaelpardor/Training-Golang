package main

import "fmt"

// An array is a numbered sequence of elements of a single type with a fixed length.
func main() {
	var x [10]int // We create an array like this
	x[9] = 777

	fmt.Println(x)                              // We print all the array.
	fmt.Println("Lenght of the array:", len(x)) // We print the lenght of the array
	fmt.Printf("Type-%T\n", x)
	fmt.Println(x[9])
}
