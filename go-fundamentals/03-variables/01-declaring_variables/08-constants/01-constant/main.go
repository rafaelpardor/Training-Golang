package main

import "fmt"

const name = "Rafael Pardo"

func main() {
	fmt.Println("We can declare global and/or constants.")
	const age = 20

	fmt.Printf("name - %v\n", name)
	fmt.Printf("age - %v\n", age)
}
