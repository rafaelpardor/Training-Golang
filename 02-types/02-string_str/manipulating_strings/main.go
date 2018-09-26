package main

import (
	"fmt"
	"strings"
)

func main() {
	getStrings()
}

func getStrings() {
	x := "Golang"
	text := "Hello world, Hello Platzi, Hello Go"

	fmt.Printf("Length of '%v' is %d bits. \n", x, len(x))
	fmt.Printf("Length of '%v' is %d bits. \n", text, len(text))
	fmt.Println(" ")
	fmt.Printf("Print the bit of the character %d. \n ", x[0])
	fmt.Printf("Print the bit of the character %d. \n ", text[0])
	fmt.Println(" ")
	fmt.Printf("Print the character %q. \n", string(x[0]))
	fmt.Printf("Print the character %q. \n", string(text[0]))
	fmt.Println(" ")
	fmt.Println(x, "/", x[0:2])
	fmt.Println(text, "/", text[0:5])
	fmt.Println(" ")
	fmt.Println(strings.Replace(x, "Go", "Og", 1))
	fmt.Println(strings.Replace(text, "Hello", "Golang", -1))
	fmt.Println(" ")
	fmt.Println(strings.ToUpper(x))
	fmt.Println(strings.ToUpper(text))
	fmt.Println(" ")
	fmt.Println(strings.ToLower(x))
	fmt.Println(strings.ToLower(text))
	fmt.Println(" ")
	fmt.Println(strings.Split(x, "G"))
	fmt.Println(strings.Split(text, "H"))
	fmt.Println(" ")
	fmt.Println("Two strings are concatenated into a single string ::", "Hello,"+" world!")
}
