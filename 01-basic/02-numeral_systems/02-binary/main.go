package main

import "fmt"

func main() {
	fmt.Println("Decimal \t Binary")
	fmt.Println("------- \t ------")
	for i := 1; i < 11; i++ {
		fmt.Printf("%d \t\t %b \n", i, i)
	}
}
