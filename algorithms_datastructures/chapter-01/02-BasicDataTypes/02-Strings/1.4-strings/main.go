package main

import "fmt"

// Stirngs are unmutables
func main() {
	s := "hello, World!"
	fmt.Println("Old string ::", s)

	r := []rune(s)
	r[0] = 'H'

	s2 := string(r)
	fmt.Println("New string ::", s2)
}
