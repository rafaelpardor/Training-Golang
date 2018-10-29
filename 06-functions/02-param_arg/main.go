package main

import "fmt"

func main() {
	// when calling greet, pass in an argument
	greet("Rafael")
	greet("Martha")
}

// greet is declared with a parameter
func greet(name string) {
	fmt.Printf("Hola, %s ¿Cómo estas?\n", name)
}
