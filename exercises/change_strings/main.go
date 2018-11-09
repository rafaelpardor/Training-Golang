package main

import "fmt"

func main() {
	stringManipulating()
}

func stringManipulating() {
	fmt.Print("\n")

	s := "hello, World!"
	fmt.Println(s)

	r := []rune(s)
	fmt.Println(r)

	r[0] = 'H'
	fmt.Println(r)
	s2 := string(r)

	fmt.Printf("New string '%s'.\n", s2)
}
