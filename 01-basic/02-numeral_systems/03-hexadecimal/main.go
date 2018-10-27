package main

import "fmt"

func main() {
	fmt.Println("Hexadecimal") // Printing hexadecimal numbers.
	fmt.Println("-----------")

	for i := 0; i < 16; i++ {
		fmt.Printf("%#X - %x \n", i, i)
	}
}
