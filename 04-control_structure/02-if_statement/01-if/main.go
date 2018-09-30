package main

import "fmt"

func main() {
	fmt.Println(1, 2)
	fmt.Println(2, 1)
}
func more(x, y int) bool {
	if x > y {
		return true
	}
	return false

}
