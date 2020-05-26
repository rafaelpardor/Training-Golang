package main

import "fmt"

// SumArray return the sum of all the elements of the integer list.
// Given list as an input argument.
func SumArray(data []int) int {
	size := len(data)
	total := 0

	for i := 0; i < size; i++ {
		total += data[i]
	}
	return total
}

// SumArrayRange return the sum of all the elements of the integer list.
// Given list as an input argument.
func SumArrayRange(data ...int) int {
	total := 0

	for _, val := range data {
		total += val
	}
	return total
}

func main() {
	xSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(SumArray(xSlice))
	fmt.Println(SumArrayRange(xSlice...))
}
