package main

import "fmt"

func foo() {
	defer fmt.Println("Yes, we are")
	defer fmt.Println("Are we done?")
	fmt.Println("Holi uwu")
}

func main() {
	foo()
}
