package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
