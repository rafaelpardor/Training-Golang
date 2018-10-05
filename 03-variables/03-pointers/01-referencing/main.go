package main

import (
	"fmt"
)

// We can save the memory adress in a variable.
func main() {
	var a int16 = 123
	var b = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Printf("Type %T\n", b)
}
