package main

import "fmt"

// We can declare a variable with a type of value, and without the respective value.
// In case that we don't declared a value on the variable, this will be have a value by default.
func main() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%q \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
