package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("The variable we have been created is a %T, with a value of %v.\n", i, i)
	}
}
