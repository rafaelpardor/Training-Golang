package main

import (
	"fmt"
)

func main() {
	for i := 1.10; i < 11.12; i++ {
		fmt.Printf("The variable we have been created, is a %T with a value of %#v.\n", i, i)
	}
}
