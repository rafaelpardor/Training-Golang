package main

import (
	"fmt"
)

func main() {
	fmt.Println("By default, a number variable is a int8.")

	for i := 1; i < 11; i++ {
		fmt.Printf("Thr variable we have been created, is a %T with a value of %#v.\n", i, i)
	}
}
