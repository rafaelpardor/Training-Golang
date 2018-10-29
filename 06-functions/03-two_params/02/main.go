package main

import "fmt"

func main() {
	greet("Rafael", "Pardo")
}

func greet(fname, lname string) {
	fmt.Println("Hola mi nombre es ", fname, lname)
}
