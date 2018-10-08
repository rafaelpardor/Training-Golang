package main

import "fmt"

func main() {
	str := "Rafael"

	for i, c := range str {
		fmt.Printf("[%d , %s]\n", i, string(c))
	}
}
