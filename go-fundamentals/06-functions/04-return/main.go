package main

import "fmt"

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func main() {
	fmt.Println("In Go, when we want to return any type of data, we need to declare it.")
	printName := greet("Rafael ", "Pardo")
	fmt.Println(printName)
}
