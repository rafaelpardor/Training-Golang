package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("->", i)
		i++

		if i%2 == 0 {
			fmt.Println(i, "Is module")
			continue
		}
		fmt.Println(">-", i)

		if i > 50 {
			break
		}
	}
}
