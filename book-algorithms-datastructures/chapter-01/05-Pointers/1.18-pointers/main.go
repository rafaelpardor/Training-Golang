package main

import "fmt"

func main() {
	data := 10
	ptr := &data

	fmt.Println("Value stored at variable data is ", data)
	fmt.Println("Value stored at variable ptr is ", *ptr)
	fmt.Println("The address of variable data is ", &data)
	fmt.Println("The address of variable ptr is ", ptr)
}
