package main

import "fmt"

// Printing decimal and binary numbers. (1-10)	(1-1010)
func main() {
	fmt.Println("Decimal \t Binary")
	fmt.Println("------- \t ------")
	for i := 1; i < 11; i++ {
		fmt.Printf("%d \t\t %b \n", i, i)
	}
}
