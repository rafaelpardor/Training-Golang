package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("\n-----------------")
	fmt.Println(mySlice)                                //We print the slice
	fmt.Println("Lenght of the slice:", len(mySlice))   //We print the lenght of the slice
	fmt.Println("Capacity of the slice:", cap(mySlice)) //We print the capacity of the slice
	fmt.Print("-----------------\n\n")

	for i := 0; i < 20; i++ {
		// Append add alements into the end of the slice
		mySlice = append(mySlice, i) //We append all the iteration that we are making
		fmt.Printf("Len: %d | Capacity: %d | Value: %d\n", len(mySlice), cap(mySlice), mySlice[i])
	}
	fmt.Printf("Len: %d | Capacity: %d\n", len(mySlice), cap(mySlice))
	fmt.Println("\n", mySlice)
}
