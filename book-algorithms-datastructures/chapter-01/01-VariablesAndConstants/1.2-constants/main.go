package main

import "fmt"

const constant1 int = 10

func main() {
	const favoriteAnimal string = "dog"

	fmt.Printf("Value stored in variable constant :: %d.\n", constant1)
	fmt.Printf("My favorite animal is a %s.\n", favoriteAnimal)
}
