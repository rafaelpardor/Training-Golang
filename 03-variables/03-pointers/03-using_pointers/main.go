package main

import "fmt"

func main() {
	a := 21
	var b = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
	*b = 123
	fmt.Println(a)
}
