package main

import (
	"fmt"
)

func main() {
	x := make([]float64, 5, 10)
	x[2] = 1.21
	x[3] = 5.521

	fmt.Println(x)
	fmt.Println("Lenght of the slice:", len(x)) // We print the lenght of the slice
	fmt.Println("Capaity of the slice:", cap(x))
	fmt.Printf("Type-%T\n", x)
}
