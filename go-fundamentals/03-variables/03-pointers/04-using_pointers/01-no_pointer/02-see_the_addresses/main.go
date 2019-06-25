package main

import "fmt"

func zero(z int) {
	z = 100
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	fmt.Println(z)
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)

	zero(x)
	fmt.Println(x) // x is still 5
}
