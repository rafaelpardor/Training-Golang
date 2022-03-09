package main

import "fmt"

func greet(fname, lname string) {
	fmt.Printf("Hi, my name is %s %s", fname, lname)
}

func main() {
	fmt.Println("In Go, when we create any numbers of parameters, we can infer the data type of the variable.")
	greet("Rafael", "Pardo")
}
