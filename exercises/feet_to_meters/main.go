package main

import "fmt"

func convertion() {
	fmt.Print("\n")

	var ft float32
	mts := ft * 0.3048

	fmt.Print("Ingresa la cantidad de feets que quieres convertir a metros: ")
	fmt.Scanf("%f", &ft)
	fmt.Println(mts, "mts.")
}

func main() {
	convertion()
}
