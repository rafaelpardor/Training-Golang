package main

import "fmt"

func max(x int) int {
	return 10 + x
}

func main() {
	sum := max(20)
	fmt.Println(sum)
}

// don't do this; bad coding practice to shadow variables
