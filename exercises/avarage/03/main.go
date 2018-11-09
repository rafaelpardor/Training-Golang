package main

import (
	"fmt"
)

func main() {
	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var total float64

	/*i, represent the current position in the array
	and value is equal at x[i] or len(x)*/
	for i, value := range x {
		fmt.Println(i, value)
		total += float64(value)
		fmt.Println(total)
	}

	fmt.Println(total)
	fmt.Println(total / float64(len(x)))
}
