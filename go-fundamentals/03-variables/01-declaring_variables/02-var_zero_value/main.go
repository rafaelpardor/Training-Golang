package main

import "fmt"

func main() {
	fmt.Println("We can declare variables with empty values, every data type has it's own default value.")

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%q \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
