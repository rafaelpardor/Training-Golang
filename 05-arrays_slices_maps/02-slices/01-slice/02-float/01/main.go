package main

import "fmt"

func main() {
	x := []float64{
		1.1,
		3.3,
		5.5,
		9.12,
		6.1,
	}

	fmt.Println(x)
	fmt.Println("Lenght of the slice:", len(x)) // We print the lenght of the slice
	fmt.Printf("Type-%T\n", x)                  // Type
}
