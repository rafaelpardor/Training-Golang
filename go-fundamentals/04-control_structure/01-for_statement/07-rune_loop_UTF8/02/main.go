package main

import "fmt"

func main() {
	for i := 48; i <= 127; i++ {
		fmt.Printf("%v \t %v \t %v \n", i, string(i), []byte(string(i)))
	}
}
