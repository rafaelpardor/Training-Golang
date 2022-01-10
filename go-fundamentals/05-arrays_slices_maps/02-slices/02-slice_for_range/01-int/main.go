package main

import "fmt"

func main() {
	xs := []int{
		1,
		3,
		5,
		7,
		9,
		11,
	}

	fmt.Println("Index \t| Value")
	for i, v := range xs {
		fmt.Printf("%d \t| %d\n", i, v)
	}
}
