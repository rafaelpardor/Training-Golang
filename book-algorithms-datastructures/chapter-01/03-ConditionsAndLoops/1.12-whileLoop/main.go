package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0

	for i < len(numbers) {
		sum += numbers[i]
		i++
	}
	fmt.Println("Sum of all numbers is ::", sum)
}
