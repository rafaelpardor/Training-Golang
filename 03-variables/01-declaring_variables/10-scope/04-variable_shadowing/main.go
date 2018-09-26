package main

import "fmt"

func max(x int) int {
	return 10 + x
}

func main() {
	max := max(20)
	fmt.Println(max)
}

// don't do this; bad coding practice to shadow variables
