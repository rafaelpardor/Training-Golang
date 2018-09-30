package main

import "fmt"

func main() {
	switch "Martha" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Juan":
		fmt.Println("Wassup Juan")
	case "Alejandro":
		fmt.Println("Wassup Alejandro")
	default:
		fmt.Println("Have you no friends?")
	}
}
