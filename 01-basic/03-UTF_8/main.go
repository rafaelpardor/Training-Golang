package main

import "fmt"

func main() {
	fmt.Print("\n")

	fmt.Println("Digital \t Binary \t Hexa \t String")
	fmt.Println("------- \t ------ \t ---- \t ------")

	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t\t %b \t %x \t %q \n", i, i, i, i)
	}
}
