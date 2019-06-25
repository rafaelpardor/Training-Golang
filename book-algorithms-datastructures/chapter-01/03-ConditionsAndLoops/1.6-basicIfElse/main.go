package main

import "fmt"

func max(x, y int) int {
	var max int

	if x > y {
		max = x
	} else {
		max = y
	}
	return max
}

func main() {
	fmt.Println(max(10, 45))
}
