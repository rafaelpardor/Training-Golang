package main

import (
	"fmt"
)

func main() {
	fmt.Println("Byte is just an alias of uint8")

	var x byte

	fmt.Printf("The variale x is a %T with a value of %v.\n", x, x)
}
