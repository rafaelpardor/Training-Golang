package main

import "fmt"

func main() {

}

func sliceAvarage() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum int

	for i := 0; i < len(number); i++ {
		sum += number[i]
		fmt.Println(sum)
	}

	fmt.Println(number)
	fmt.Println(sum)
	fmt.Println("Sum is ::", sum)
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "ds", "42", "sagsf")

	fmt.Println(slice1)
}
