package main

import (
	"fmt"
)

func main() {
	basicExmple()
}

func basicExmple() {
	fmt.Print("\n")

	var input float64
	output := input * 3

	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &input)

	fmt.Printf("%d times 3 is equal to: %d\n", int32(input), int32(output))
}
