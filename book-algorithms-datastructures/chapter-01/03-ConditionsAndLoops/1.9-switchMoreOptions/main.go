package main

import (
	"fmt"
)

func main() {
	i := 2

	switch i {
	case 1, 2, 3:
		fmt.Println("One, Two or Three")
	default:
		fmt.Println("Something else")
	}
}
