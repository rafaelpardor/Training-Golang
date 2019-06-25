package main

import (
	"fmt"
)

func main() {
	x := make([]int, 5, 10)
	x[1] = 1
	x[4] = 5

	fmt.Println(x)
	fmt.Println("Lenght of the slice:", len(x))
	fmt.Println("Capaity of the slice:", cap(x))
	fmt.Printf("Type-%T\n", x)
}
