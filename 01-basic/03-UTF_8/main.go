package main

import "fmt"

func main() {
	fmt.Println("Digital \t Binary \t Hexa \t String")
	fmt.Println("------- \t ------ \t ---- \t ------")
	for i := 32; i < 127; i++ {
		fmt.Printf("%d \t\t %b \t %x \t %q \n", i, i, i, i)
	}
}
