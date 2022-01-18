package main

import "fmt"

func printSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(s)

	// [1 2 3 4 5 6 7 8 9 10] :: len=10 cap=10
	a := make([]int, 10)
	printSlice(a)

	// [] :: len=0 cap=10
	c := s[0:4]
	printSlice(c) // [1 2 3 4] :: len=4 cap=10

	d := c[2:5]
	printSlice(d) // [3 4 5] :: len=3 cap=8
}
