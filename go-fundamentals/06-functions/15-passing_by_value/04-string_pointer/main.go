package main

import "fmt"

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z) // Todd
	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z) // Rocky
}

func main() {
	name := "Todd"
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println(&name)
	fmt.Println(name) //Rocky
}
