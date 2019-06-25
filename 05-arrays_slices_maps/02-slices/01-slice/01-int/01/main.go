package main

import "fmt"

func main() {
	x := []int{
		1,
		3,
		5,
		9,
		11,
	}

	fmt.Println(x)
	fmt.Println("Lenght of the slice:", len(x))
	fmt.Printf("Type-%T\n", x)
}
