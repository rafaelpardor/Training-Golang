package main

import (
	"fmt"
)

func main() {
	fmt.Println("-.- Byte type -.-")
	funcByte()
}
func funcByte() {
	fmt.Print("\n")

	fmt.Println("Byte is just an alias of uint8")

	var v byte

	fmt.Printf("%T %v.\n", v, v)
}
