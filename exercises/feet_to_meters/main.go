package main

import "fmt"

var feets float32

func main() {
	fmt.Print("Ingresa la cantidad de feets que quieres convertir a metros: ")

	fmt.Scan(&feets)
	mts := feets * 0.3048

	fmt.Println(mts, "mts.")
}
