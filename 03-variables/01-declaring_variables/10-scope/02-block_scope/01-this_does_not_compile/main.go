package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	// foo() will not work
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	fmt.Println(x)
}
