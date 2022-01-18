package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	fmt.Printf("Length of array :: %v.\n", len(arr))
}
