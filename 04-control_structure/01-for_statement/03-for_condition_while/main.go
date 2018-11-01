package main

import "fmt"

func main() {
	// for <condition> { } - like a while loop
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
