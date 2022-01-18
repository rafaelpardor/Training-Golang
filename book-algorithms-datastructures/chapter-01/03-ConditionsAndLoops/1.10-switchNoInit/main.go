package main

import "fmt"

func isEven(value int) {
	switch {
	case value%2 == 0:
		fmt.Println("value is even")
	default:
		fmt.Println("value is odd")
	}
}

func main() {
	isEven(10)
}
