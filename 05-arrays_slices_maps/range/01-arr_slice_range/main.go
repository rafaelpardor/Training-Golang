package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\n", numbers)
	sum := 0

	fmt.Println("Index , Valor")
	for i, val := range numbers {
		sum += val
		fmt.Println("[", i, ",", val, "] ")
	}
	fmt.Println("Sum is :: ", sum)
}
