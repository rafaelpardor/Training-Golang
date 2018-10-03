package main

import (
	"fmt"
	"strings"
)

func main() {
	getStrings()
}

func getStrings() {
	aa := "Golang"
	ab := "Hello world, Hello Platzi, Hello Go"

	lengthOfStrings(aa)
	lengthOfStrings(ab)

	fmt.Println(" ")
	decimalString(aa)
	decimalString(ab)

	// todo Pass to function
	fmt.Println(" ")
	fmt.Println(aa, "/", aa[0:1])
	fmt.Println(ab, "/", ab[0:5])

	fmt.Println(" ")
	fmt.Println(strings.Replace(aa, "Go", "Og", 1))
	fmt.Println(strings.Replace(ab, "Hello", "Golang", -1))
	fmt.Println(" ")
	fmt.Println(strings.ToUpper(aa))
	fmt.Println(strings.ToUpper(ab))
	fmt.Println(" ")
	fmt.Println(strings.ToLower(aa))
	fmt.Println(strings.ToLower(ab))
	fmt.Println(" ")
	fmt.Println(strings.Split(aa, "G"))
	fmt.Println(strings.Split(ab, "H"))
	fmt.Println(" ")
	fmt.Println("Two strings are concatenated into a single string ::", "Hello,"+" world!")
}
