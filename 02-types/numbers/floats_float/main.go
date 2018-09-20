package main

import (
	"fmt"
)

func main() {
	fmt.Println("-.- Floats -.-")
	floats()
	fmt.Println(getFloat(0.1, 0.12))
}

func floats() {
	fmt.Print("\n")

	var x float32 = 32.55345

	fmt.Printf("The variable we have been created is a %T with a value of %v\n", x, x)
}

func getFloat(a float32, b float64) (float32, float64) {
	fmt.Print("\n")

	return float32(a), float64(float32(b))
}
