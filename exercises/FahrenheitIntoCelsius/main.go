package main

import "fmt"

func main() {
	convertion()
}

func convertion() {
	fmt.Print("\n")

	var F float32
	C := ((F - 32) * 5 / 9)

	fmt.Print("Ingresa la cantidad de grados Fahrenheit que quieres convertir a Celcius: ")
	fmt.Scanf("%f", &F)

	fmt.Printf("%f degrees \n", C)
}
