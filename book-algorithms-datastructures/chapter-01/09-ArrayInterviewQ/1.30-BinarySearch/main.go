package main

import "fmt"

// BinarySearch in a sorted list
// Binary search is an efficient algorithm for finding an item from a sorted list of items.
// It works by repeatedly dividing in half the portion of the list that could contain the item,
// until you've narrowed down the possible locations to just one.
func BinarySearch(data []int, value int) bool {
	size := len(data)
	var mid int
	low := 0
	high := size - 1
	fmt.Printf("Len: %v, mid: %v, low: %v, high:  %v\n", size, mid, low, high)

	for low <= high {
		mid = low + (high-low)/2
		fmt.Printf("mid value: %v\n", mid)

		if data[mid] != value {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			return true
		}
	}
	return false
}

func main() {
	xSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// array of numbers / number to guess
	fmt.Println(BinarySearch(xSlice, 50))
}
