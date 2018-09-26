package main

import (
	"fmt"
)

func main() {
	fmt.Println("Byte is just an alias of uint8")

	var x byte

	fmt.Printf("%T %v.\n", x, x)
}
