package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println("Lenght of the slice:", len(mySlice))
	fmt.Println("Capacity of the slice:", cap(mySlice))
	fmt.Print("-----------------\n\n")

	for i := 0; i < 20; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len: %d | Capacity: %d | Value: %d\n", len(mySlice), cap(mySlice), mySlice[i])
	}
	fmt.Printf("\nLen: %d | Capacity: %d\n", len(mySlice), cap(mySlice))
	fmt.Println(mySlice)
}
