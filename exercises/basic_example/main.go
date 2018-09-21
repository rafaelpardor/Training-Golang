package main

import (
	"fmt"
)

func main() {
	basicExmple()
}

func basicExmple() {
	fmt.Print("\n")

	var input int16
	var times int16

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &input)
	fmt.Print("Enter another number: ")
	fmt.Scanf("%d", &times)

	output := input * times

	fmt.Printf("%d times %d is equal to: %d\n", input, times, output)
}
