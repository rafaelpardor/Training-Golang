package main

import (
	"fmt"
)

var name string

func main() {
	fmt.Print("Please write your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello your name is, %s, welcome to Go!.\n", name)
}
