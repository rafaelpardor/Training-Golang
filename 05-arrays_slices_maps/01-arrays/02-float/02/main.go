package main

import (
	"fmt"
)

func main() {
	x := [6]float32{98.1, 93.54, 7.21, 2.76, 1.2}

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x)) // We print the lenght of the array
	fmt.Printf("Type-%T\n", x)
}
