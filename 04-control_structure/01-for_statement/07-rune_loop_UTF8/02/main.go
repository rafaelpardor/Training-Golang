package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v \t %v \t %v \n", i, string(i), []byte(string(i)))
	}
}
