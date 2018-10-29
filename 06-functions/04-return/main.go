package main

import "fmt"

func main() {
	fmt.Println(greet("Rafael ", "Pardo"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
