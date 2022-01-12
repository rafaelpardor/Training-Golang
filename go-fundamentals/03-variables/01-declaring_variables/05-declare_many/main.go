package main

import "fmt"

func main() {
	fmt.Println("We can declare multiple variables in the same line.")
	var message string
	var a, b, c int

	a = 1
	// b = 0 by default
	// c = 0 by default
	message = "Hello world!"

	fmt.Println(message, a, b, c)
}
