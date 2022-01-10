package main

import (
	"fmt"
)

func main() {
	fmt.Println("Byte is just an alias of uint8, this meaning that it will only receive positive numbers.")

	var x byte = 10
	//	var err byte = -10
	byteArr := []byte{1, 2, 3}

	fmt.Printf("The variale x is a '<%T>' with a value of %v.\n", x, x)
	fmt.Printf("The variale byteArr is a '<%T>' with a value of %v.\n", byteArr, byteArr)
}
