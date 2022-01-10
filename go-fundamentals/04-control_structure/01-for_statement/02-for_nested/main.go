package main

import "fmt"

func main() {
	for i := 1; i < 6; i++ {
		fmt.Printf("%v\n", i)
		for j := 1; j < 6; j++ {
			fmt.Printf("\t%v\n", j)
			for k := 1; k < 6; k++ {
				fmt.Printf("\t\t%v\n", k)
			}
		}
	}
}
