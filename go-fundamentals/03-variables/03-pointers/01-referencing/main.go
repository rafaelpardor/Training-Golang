package main

import (
	"fmt"
)

func main() {
	fmt.Println("In Go we can save the address memory in a variable, this is called referencing.")
	a := 123
	b := &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Printf("Type %T\n", b)
}
