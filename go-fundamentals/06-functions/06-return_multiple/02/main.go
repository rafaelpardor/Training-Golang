package main

import "fmt"

func greet(fname, lname string) (string, string) {
	return fmt.Sprintln("Hola soy,", fname, lname), fmt.Sprintln(lname, fname)
}

func main() {
	fmt.Print(greet("Rafael", "Pardo"))
}
