package main

import "fmt"

func main() {
	// x undefined
	fmt.Println(x)
	fmt.Println(y)

	x := 42
}

var y = 42
