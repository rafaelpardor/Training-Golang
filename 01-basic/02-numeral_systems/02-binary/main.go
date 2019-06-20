package main

import "fmt"

func main() {
	fmt.Println("Binary numbers")
	fmt.Println("-------")

	for i := 0; i < 10; i++ {
		fmt.Printf("%d - %b \n", i, i)
	}
}
