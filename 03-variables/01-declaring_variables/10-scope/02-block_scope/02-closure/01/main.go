package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	// Doing this, we create a type of block, with his own scope.
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}
	{
		x := 100
		fmt.Println(x)
	}
	fmt.Println(x)
}
