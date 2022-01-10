package main

import "fmt"

type foo int
type bar string

func main() {
	var myName bar = "Rafael" // var <name> <foo = int>
	var myAge foo

	myAge = 20
	fmt.Printf("%T %v\n", myName, myName)
	fmt.Printf("%T %v\n", &myName, myName)
	fmt.Printf("%T %v\n", myAge, myAge)
	fmt.Printf("%T %v\n", &myAge, myAge)
}
