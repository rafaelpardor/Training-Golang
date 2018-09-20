package main

import "fmt"

func main() {
	fmt.Print("\n")

	fmt.Printf("Decimal \t Binary \t\t Hexadecimal\n")
	fmt.Println("------- \t -------------------- \t -----------")

	for i := 1000000; i <= 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
