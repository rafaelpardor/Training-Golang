package main

import "fmt"

// We can declared multiple variables in one same line.
// In this case, a,b and c are the same type of value, and, b and c have no value declared.
func main() {
	var message string
	var a, b, c int

	a = 1
	// b = 0 int by default
	// c = 0 int by default
	message = "Hello world!"

	fmt.Println(message, a, b, c)
}
