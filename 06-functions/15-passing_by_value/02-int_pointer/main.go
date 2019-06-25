package main

import "fmt"

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 2
	fmt.Println(z)
	fmt.Println(*z)
}

func main() {
	age := 44
	fmt.Println(age)
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(age)
	fmt.Println(&age)
}
