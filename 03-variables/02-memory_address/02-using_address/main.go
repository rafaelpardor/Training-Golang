package main

import "fmt"

const metersToYards float64 = 1.09361

// We can manipulate a variable using the memory adress
func main() {
	var meters float64

	fmt.Print("Enter meters swam: ")
	fmt.Scanf("%f", &meters)

	yards := meters * metersToYards

	fmt.Println(meters, " meters is ", yards, " yards.")
}
