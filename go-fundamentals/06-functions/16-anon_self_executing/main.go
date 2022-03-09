package main

import "fmt"

func main() {
	fmt.Println("In Go, we can create anonymous functions")
	func() {
		fmt.Println("Hello world")
	}()
}
