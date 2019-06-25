package main

import "fmt"

func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0

	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func main() {
	data := []int{1, 2}

	fmt.Println("Max sub array sum :", MaxSubArraySum(data))
}
