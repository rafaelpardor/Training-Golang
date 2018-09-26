package main

import "fmt"

var x = 19

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
