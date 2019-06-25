package main

import "fmt"

func main() {
	xs := []string{
		"Rafael",
		"Goalng",
		"Google",
		"Martha",
		"Pardo",
	}

	fmt.Println("Index \t| Value")
	for i, v := range xs {
		fmt.Printf("%d \t| %s\n", i, v)
	}
}
