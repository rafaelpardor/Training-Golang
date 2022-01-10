package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	print("\n")
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	print("\n")
}

func main() {
	fmt.Println("A string is a slice of bytes, they can be created by enclosing a set of characters inside double quotes.")

	var myName string
	myName = "Rafael Pardo"

	fmt.Printf("The variable myName is a '%T', type with the value %v\n", myName, myName)
	printChars(myName)
	printBytes(myName)
}
