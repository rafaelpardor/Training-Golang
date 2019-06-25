package main

import (
	"fmt"
)

func main() {
	x, y, z := f()
	fmt.Println(x, y, z)
}

func f() (int, int, int) {
	return 5, 6, 10
}
