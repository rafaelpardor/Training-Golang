package main

import "fmt"

func operation(sum int) (x, y, z, a int, b, c string) {
	x = sum * 4 / 9
	y = sum - x
	z = sum * sum
	a = x + y + sum
	b = "Rafael"
	c = "Edmonton"

	return
}

func main() {
	fmt.Println(operation(10))
}
