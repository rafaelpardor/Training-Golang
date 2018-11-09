package main

import "fmt"

func main() {
	if x := 2; x%2 == 0 {
		fmt.Println(x, "even")
	} else {
		fmt.Println(x, "odd")
	}
}
