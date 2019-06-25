package main

import "fmt"

func main() {
	var x [10]int

	fmt.Println(x)
	fmt.Println("Lenght of the array:", len(x))
	fmt.Println("Capaity of the slice:", cap(x))
	fmt.Printf("Type-%T\n", x)

	x[9] = 777
	fmt.Println(x[9])
	fmt.Println(x)
}
