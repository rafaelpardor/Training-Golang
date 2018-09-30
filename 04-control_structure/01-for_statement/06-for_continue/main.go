package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println(i)
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 50 {
			break
		}
	}
}
