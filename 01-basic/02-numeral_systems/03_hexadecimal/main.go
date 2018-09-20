package main

import "fmt"

func main() {
	x := 42
	fmt.Println("Decimal \t Binary \t Hexadecimal")
	fmt.Println("------- \t ------ \t -----------")
	fmt.Printf("%d \t\t %b \t %#X-%x \n", x, x, x, x)

}
