package main

import "fmt"

var f float32

func main() {
	fmt.Print("Ingresa la cantidad de grados Fahrenheit que quieres convertir a Celcius: ")

	fmt.Scanf("%f", &f)
	c := ((f - 32) * 5 / 9)

	fmt.Printf("%f degrees.\n", c)
}
