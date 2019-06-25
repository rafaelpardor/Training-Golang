package main

import "fmt"

func changeMe(z string) {
	fmt.Println(z) // Todd
	z = "Rocky"
	fmt.Println(z) // Rocky
}

func main() {
	name := "Todd"
	fmt.Println(name) // Todd

	changeMe(name)

	fmt.Println(name) // Todd
}
