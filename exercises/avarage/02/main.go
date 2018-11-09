package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
