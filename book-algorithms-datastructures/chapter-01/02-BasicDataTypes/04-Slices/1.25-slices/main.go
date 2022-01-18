package main

import "fmt"

// printSlice take an slice as argument
// and show basics properties.
func printSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}

func main() {
	var s []int

	for i := 1; i <= 17; i++ {
		s = append(s, i)
		printSlice(s)
	}
}
