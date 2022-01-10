package main

import "fmt"

func main() {
	a := 21
	b := &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
	*b = 123
	fmt.Println(a)
}
