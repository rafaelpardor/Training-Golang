package main

import "fmt"

// An array is a numbered sequence of elements of a single type with a fixed length.
func main() {
	var x [10]string // We create an array like this
	x[0] = "-"
	x[1] = "Rafael"
	x[5] = "Martha"

	fmt.Println(x)                              // We print all the array.
	fmt.Println("Lenght of the array:", len(x)) // We print the lenght of the array
	fmt.Printf("Type-%T\n", x)
}
