package main

import "fmt"

// The for statement allows us to repeat a list of statements (a block) multiple times.
func main() {
	// for <initialization>; <condition>; <increment/decrement>
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
