package main

import "fmt"

func main() {
	/*
		The syntax of switch for type.
		switch var.(type){
		case <type>:
			<statements>
		case <type>:
			<statements>
		default:
			<statements>
		}
	*/
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
