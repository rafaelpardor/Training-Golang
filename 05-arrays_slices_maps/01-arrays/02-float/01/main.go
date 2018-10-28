package main

import "fmt"

// An array is a numbered sequence of elements of a single type with a fixed length.
func main() {
	var x [10]float32 // We create an array like this
	x[1] = 777.1
	x[2] = 2.1483

	fmt.Println(x)                              // We print all the array.
	fmt.Println("Lenght of the array:", len(x)) // We print the lenght of the array
	fmt.Printf("Type-%T\n", x)
}
