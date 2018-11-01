package main

import "fmt"

func main() {
	fmt.Println("Hexadecimal") // Printing hexadecimal numbers.
	fmt.Println("-----------")

	for i := 0; i < 16; i++ {
		fmt.Printf("%#X - %x \n", i, i) // Printing hexadecimal, 0 to 9 and a to f.
	}
}
