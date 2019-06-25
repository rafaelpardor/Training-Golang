package main

import "fmt"

func greet(fname, lname string) {
	fmt.Printf("Hola mi nombre es %s %s", fname, lname)
}

func main() {
	greet("Rafael", "Pardo")
}
