package main

import (
	"fmt"
)

func main() {
	fmt.Println("-.- Rune -.-")
	funcRune()
}

func funcRune() {
	fmt.Print("\n")

	fmt.Print("What rune does is descompose any string, and return an array of his respectives bits charactes.\n\n")

	x := "This is a all string changed to an array of integers."
	newArray := []rune(x)

	fmt.Println(x)
	fmt.Println(newArray)
}
