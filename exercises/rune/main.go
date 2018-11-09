package main

import (
	"fmt"
)

//What rune does is descompose any string, and return an slice of his respectives decimal charactes.
func main() {
	x := "This is a all string changed to an array of integers."
	newArray := []rune(x)

	fmt.Println(x)
	fmt.Println(newArray)
	fmt.Println(len(newArray))
}
