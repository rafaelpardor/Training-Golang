package main

import (
	"fmt"
)

func main() {
	fmt.Println("The characteristic of the 'uint' type is that only saves positive numbers.")
	var c uint = 121
	// var err uint = -1

	fmt.Printf("Type: '<%T>', value: %v\n", c, c)
	// fmt.Println(err)
}
