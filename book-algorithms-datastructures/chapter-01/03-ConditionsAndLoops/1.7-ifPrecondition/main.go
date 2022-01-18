package main

import "fmt"

func maxAreaCheck(length, width, limit int) bool {
	if area := length * width; area < limit {
		return true
	}
	return false
}

func main() {
	fmt.Println(maxAreaCheck(10, 15, 4))
}
