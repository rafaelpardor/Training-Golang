package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}
	myAnotherSlice := []int{10, 11, 12, 13}

	mySlice = append(mySlice, myOtherSlice...)
	mySlice = append(mySlice, myAnotherSlice...)
	mySlice = append(mySlice, 14, 15)

	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))
}
