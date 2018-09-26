package main

import (
	"fmt"
)

func main() {
	fmt.Println("By default, a float variable is a float64.")

	for i := 1.123; i < 11.123; i++ {
		fmt.Printf("Thr variable we have been created, is a %T with a value of %#v.\n", i, i)
	}
}
