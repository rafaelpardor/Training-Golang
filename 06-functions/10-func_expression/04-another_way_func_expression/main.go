package main

import "fmt"

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}
