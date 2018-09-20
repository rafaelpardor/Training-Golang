package main

import "fmt"

func main() {
	fmt.Print("\n")

	x := 42
	fmt.Println("Decimal \t Binary \t Hexadecimal")
	fmt.Println("------- \t ------ \t -----------")
	fmt.Printf("%d \t\t %b \t %#X-%x \n", x, x, x, x)

}
