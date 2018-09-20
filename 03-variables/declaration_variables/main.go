package main

import "fmt"

func main() {
	fmt.Println("-.- Var declaration -.-")
	varDeclaration()
	shortCut()
}

func varDeclaration() {
	fmt.Print("\n")

	var v1 int8
	var v2 int // If a variable is declarated and this no has value, by default it will be 0
	v1 = 10
	var x = "Platzi is the best place to learn cool stuff."

	fmt.Printf("Value stored in variable v1 :: %d.\n", v1)
	fmt.Printf("Value stored in variable v2 :: %d.\n", v2)
	fmt.Println(x)
}

func shortCut() {
	fmt.Print("\n")

	fmt.Println("':=' operator allows short declaration of variable. This does declaration and initialization of variable at once. The type of the variable is automatically determined by the compiler according to the type of value assigned")

	myDogName := "Luna"
	numbers := 12 + 32*3 + (3 / 4)

	fmt.Println(myDogName)
	fmt.Println(numbers)
}
