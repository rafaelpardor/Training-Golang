package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)

	fmt.Println("\n-----------------")
	fmt.Println(mySlice)
	fmt.Println("Lenght of the slice:", len(mySlice))
	fmt.Println("Capacity of the lenght:", cap(mySlice))
	fmt.Print("-----------------\n\n")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len: %d | Capacity: %d | Value: %d\n", len(mySlice), cap(mySlice), mySlice[i])
	}
	fmt.Println("\n", mySlice)
}
