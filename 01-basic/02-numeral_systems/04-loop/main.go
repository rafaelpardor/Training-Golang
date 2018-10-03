package main

import "fmt"

func main() {
	fmt.Printf("Decimal \t Binary \t\t Hexadecimal\n")
	fmt.Println("------- \t -------------------- \t -----------")
	for i := 1; i < 256; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
