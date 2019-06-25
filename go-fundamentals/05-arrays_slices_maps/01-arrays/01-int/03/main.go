package main

import (
	"fmt"
)

func main() {
	x := [5]int{
		98,
		93,
		77,
		82,
		// 83,
	}

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x))
	fmt.Printf("Type-%T\n", x)
}
