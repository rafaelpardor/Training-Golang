package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hola, %s ¿Cómo estas?\n", name)
}

func main() {
	greet("Rafael")
	greet("Martha")
}
