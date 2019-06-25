package main

import "fmt"

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func main() {
	printName := greet("Rafael ", "Pardo")
	fmt.Println(printName)
}
