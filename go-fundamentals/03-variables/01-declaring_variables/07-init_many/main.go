package main

import "fmt"

// In this case we daclared many variables, all with one type of value and value in order of invocation.
func main() {
	var message = "Hello world!"
	var a, b, c int = 1, 2, 3

	fmt.Println(message, a, b, c)
}
