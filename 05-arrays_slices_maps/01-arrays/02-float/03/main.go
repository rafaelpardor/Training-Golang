package main

import "fmt"

func main() {
	x := [5]float64{
		98.32,
		93.87,
		77.1111,
		82.12,
		// 83.876,
	}

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x))
	fmt.Printf("Type-%T\n", x)
}
