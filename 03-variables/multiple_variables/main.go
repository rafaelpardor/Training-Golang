package main

import "fmt"

func main() {
	fmt.Println("-.- Multiple var declaration -.-")
	multipleVariables()
}

func multipleVariables() {
	fmt.Print("\n")

	var (
		a = 12.1
		b = 324
		c = 3 * 5
		d = "rafaelpardor"
		e = true
	)
	fmt.Println(a, b, c, d, e)
}
