package main

import "fmt"

func main() {
	fmt.Println("Decimal \t Binary \t Hexadecimal")
	fmt.Println("------- \t ------ \t -----------")
	for i := 1; i < 16; i++ {
		fmt.Printf("%d \t\t %b \t\t %#X-%x \n", i, i, i, i)
	}
}
