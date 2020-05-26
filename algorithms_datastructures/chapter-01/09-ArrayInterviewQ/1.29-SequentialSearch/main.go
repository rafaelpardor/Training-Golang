package main

import "fmt"

// SequentialSearch will search a list for som given value.
func SequentialSearch(data []int, value int) bool {
	size := len(data)

	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func main() {
	xSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(SequentialSearch(xSlice, 3))
	fmt.Println(SequentialSearch(xSlice, 9))
}
