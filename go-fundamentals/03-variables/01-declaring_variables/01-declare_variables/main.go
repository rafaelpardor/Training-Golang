package main

import "fmt"

func main() {
	fmt.Println("In Go we can declare variables with the keywork 'var', followed by the name of the variable and the data type.")

	var greeting string
	var numbers int = 100

	greeting = "Hello World."

	fmt.Println(greeting)
	fmt.Println(numbers)
}
