package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// undefined food
	fmt.Println(food)

}
