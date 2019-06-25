package main

import (
	"fmt"
	"math"
)

func main() {
	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64

	fmt.Println("The float type has two types, 'float32' and 'float64'")
	fmt.Printf("The variable we have been created, is a '<%T>' with a value of %f.\n", maxFloat32, maxFloat32)
	fmt.Printf("The variable we have been created, is a '<%T>' with a value of %3.2f.\n", maxFloat64, maxFloat64)
}
