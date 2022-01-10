package main

import (
	"fmt"
)

func main() {
	x := [6]float64{98.1, 93.54, 7.21, 2.76, 1.2}

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x))
	fmt.Printf("Type-%T\n", x)
}
