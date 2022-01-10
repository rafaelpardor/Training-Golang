package main

import "fmt"

type customer struct {
	name string
	age  int
}

func changeMe(z *customer) {
	fmt.Println(z) // &{Todd 44}
	fmt.Println(&z.name)
	z.name = "Rocky"
	fmt.Println(z) // &{Rocky 44}
	fmt.Println(&z.name)
}

func main() {
	c1 := customer{"Todd", 44}
	fmt.Println(&c1.name)

	changeMe(&c1)

	fmt.Println(c1) // {Rocky 44}
	fmt.Println(&c1.name)
}
