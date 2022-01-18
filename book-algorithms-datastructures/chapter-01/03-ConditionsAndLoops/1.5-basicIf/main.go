package main

import "fmt"

func more(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

func main() {
	fmt.Println(more(10, 4))
}
