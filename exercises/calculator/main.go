package main

import "fmt"

func main() {
	welcome()
	result()
	goodBye()
}

func welcome() {
	fmt.Println("          /*************************************\\")
	fmt.Println("********** Bienvenido a calulator v0.1 by Platzi **********")
	fmt.Println("          \\*************************************/")
}

func getNumbers() (int, int) {
	var numbert1 int
	var numbert2 int

	fmt.Println("")
	fmt.Println("Ingresa dos números ENTEROS separados por un espacio.\nEl 0 es el valor por defecto.")
	fmt.Scanf("%d %d", &numbert1, &numbert2)

	return numbert1, numbert2
}
func result() {
	numbert1, numbert2 := getNumbers()
	fmt.Printf("Número %d y número %d.\n", numbert1, numbert2)
	fmt.Printf("Sumatoria = %d.\n", sum(numbert1, numbert2))
	fmt.Printf("Resta = %d.\n", sust(numbert1, numbert2))
	fmt.Printf("Multiplicación = %d\n.", mult(numbert1, numbert2))
	fmt.Printf("División = %d.\n", div(numbert1, numbert2))
}

func goodBye() {
	fmt.Println("Hasta luego.")
}

func sum(a int, b int) int {
	return a + b
}
func sust(a int, b int) int {
	return a - b
}
func mult(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}
