package main

import (
	"fmt"
)

// We can save the memory adress in a variable.
func main() {
	a := 123
	b := &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Printf("Type %T\n", b)
}
