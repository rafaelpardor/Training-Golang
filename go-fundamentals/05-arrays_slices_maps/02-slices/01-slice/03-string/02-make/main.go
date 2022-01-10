package main

import (
	"fmt"
)

func main() {
	x := make([]string, 5, 10)
	x[1] = "Rafael"
	x[4] = "Martha"

	fmt.Println(x)
	fmt.Println("Lenght of the slice:", len(x))
	fmt.Println("Capaity of the slice:", cap(x))
	fmt.Printf("Type-%T\n", x)
}
