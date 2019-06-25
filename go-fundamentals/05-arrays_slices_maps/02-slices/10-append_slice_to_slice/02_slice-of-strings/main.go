package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday", "Wednesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}
	anotherSlice := []string{"basdf", "fsa"}

	mySlice = append(mySlice, myOtherSlice...)
	mySlice = append(mySlice, anotherSlice...)

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
