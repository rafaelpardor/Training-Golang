package main

import "fmt"

// We can print the value of the memory adress using *
func main() {
	a := 20
	var b = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)
}
