package main

import "fmt"

func main() {
	var x [10]string
	x[0] = "-"
	x[5] = "Martha"
	x[9] = "Rafael"

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x))
	fmt.Printf("Type-%T\n", x)
}
