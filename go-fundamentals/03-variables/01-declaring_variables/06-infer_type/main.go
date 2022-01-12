package main

import "fmt"

// We can declared multiple variables in one same line.
// In this case
func main() {
	fmt.Println("We can declare and assing multiple variables in the same line.")
	var message = "Hello world!"
	var a, b, c = 1, 2, 3

	fmt.Println(message, a, b, c)
}
