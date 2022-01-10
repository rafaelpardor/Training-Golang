package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 20
	fmt.Printf("%T %v \n", myAge, myAge)
	fmt.Printf("%T %v \n", int(myAge), myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// Esto no funciona:
	// fmt.Println(myAge + yourAge)

	// Esta convercion funciona:
	fmt.Println(int(myAge) + yourAge)
}
