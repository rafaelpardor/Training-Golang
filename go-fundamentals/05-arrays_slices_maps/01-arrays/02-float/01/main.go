package main

import "fmt"

func main() {
	var x [10]float64
	x[1] = 777.1
	x[5] = 2.1483

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x))
	fmt.Printf("Type-%T\n", x)
}
