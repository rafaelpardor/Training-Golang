package main

import "fmt"

func RotateArray(data []int, k int) []int {
	n := len(data)

	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)

	return data
}

func ReverseArray(data []int, start int, end int) {
	i := start
	j := end

	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func main() {
	xSlice := []int{0, 1, 2, 3}

	fmt.Println(RotateArray(xSlice, 3))
}
