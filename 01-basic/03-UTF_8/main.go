package main

import "fmt"

func main() {
	fmt.Println("UTF Codes")
	fmt.Println("-------")

	for i := 32; i < 127; i++ {
		fmt.Printf("%d ->\t %q \n", i, i)
	}
}
