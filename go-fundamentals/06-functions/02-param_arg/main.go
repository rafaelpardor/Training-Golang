package main

import "fmt"

func greet(name string) {
	fmt.Printf("My name is %s\n", name)
}

func main() {
	fmt.Println("Functions can receive parameters, in order of 'name' and 'type'")
	greet("Rafael")
	greet("Rob")
}
