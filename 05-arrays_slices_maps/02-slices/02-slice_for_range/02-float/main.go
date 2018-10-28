package main

import "fmt"

func main() {
	xs := []float64{
		1.1,
		3.3,
		5.5,
		9.12,
		6.1,
	}

	fmt.Println("Index \t| Value")
	for i, value := range xs {
		fmt.Printf("%d \t| %f\n", i, value)
	}
}
