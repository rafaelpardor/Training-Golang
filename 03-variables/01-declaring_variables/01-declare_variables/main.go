package main

import "fmt"

// We can declare a variable with a type of value, and without the respective value.
// We assigned the value, after we declared the value.
func main() {
	var greeting string
	var numbers int

	greeting = "Hello World."
	numbers = 100

	fmt.Println(greeting)
	fmt.Println(numbers)
}
