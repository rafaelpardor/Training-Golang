package main

import "fmt"

func greet(fname string, lname string) {
	fmt.Printf("Hola mi nombre es %s %s\n", fname, lname)
}

func main() {
	greet("Rafael", "Pardo")
}
