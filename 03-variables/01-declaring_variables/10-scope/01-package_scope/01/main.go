package main

import "fmt"

// This is a global variable.
var x = 19

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
