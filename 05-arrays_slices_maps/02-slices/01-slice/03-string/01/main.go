package main

import "fmt"

func main() {
	x := []string{
		"Rafael",
		"Goalng",
		"Google",
		"Martha",
		"Pardo",
	}

	fmt.Println(x)
	fmt.Println("Lenght of the slice:", len(x)) // We print the lenght of the slice
	fmt.Println("Capaity of the slice:", cap(x))
	fmt.Printf("Type-%T\n", x) // Type
}
