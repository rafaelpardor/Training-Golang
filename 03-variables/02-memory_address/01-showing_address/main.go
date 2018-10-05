package main

import "fmt"

// We can print the memory adress of a variable.
func main() {
	a := 12

	fmt.Printf("a - %v\n", a)
	fmt.Println("a variable memory address -", &a)
	fmt.Printf("%d \n", &a)
}
