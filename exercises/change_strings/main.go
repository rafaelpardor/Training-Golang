package main

import "fmt"

func main() {
	stringManipulating()
}

func stringManipulating() {
	fmt.Print("\n")

	fmt.Print("Strings are immutable so you cannot change its content once created.\n And in the end convert it back to string.\n\n")

	s := "hello, World!"
	fmt.Println(s)

	r := []rune(s)
	fmt.Println(r)

	r[0] = 'H'
	fmt.Println(r)
	s2 := string(r)

	fmt.Printf("New string '%s'.\n", s2)
}
