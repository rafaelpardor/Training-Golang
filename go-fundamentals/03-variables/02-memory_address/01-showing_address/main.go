package main

import "fmt"

func main() {
	fmt.Println("Go supports pointers, by creating a variable and usign the keywork '&', we can access the address memory.")
	a := 12

	fmt.Printf("a - %v\n", a)
	fmt.Println("a variable memory address -", &a)
}
