package main

import (
	"fmt"

	"github.com/rafaelpardor/golang/exercises/calculator/calc"
)

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
	fmt.Printf("Sumatoria = %d.\n", calc.Sum(numbert1, numbert2))
	fmt.Printf("Resta = %d.\n", calc.Sust(numbert1, numbert2))
	fmt.Printf("Multiplicación = %d\n.", calc.Mult(numbert1, numbert2))
	fmt.Printf("División = %d.\n", calc.Div(numbert1, numbert2))
}

func goodBye() {
	fmt.Println("Hasta luego.")
}
