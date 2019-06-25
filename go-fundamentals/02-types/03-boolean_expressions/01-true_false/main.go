package main

import "fmt"

func main() {
	fmt.Println("Booleans are a type of data that Go supports")
	if true {
		fmt.Println("This code will ran.")
	}

	if false {
		fmt.Println("This did not run.")
	}
}
