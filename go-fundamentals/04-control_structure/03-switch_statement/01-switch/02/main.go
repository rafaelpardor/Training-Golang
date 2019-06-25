package main

import "fmt"

func main() {
	/*
		The syntax of switch with conditions
		switch {
		case <condition>:
			<statements>
		case <condition>:
			<statements>
		default:
			<statements>
		}
	*/
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("something else")
	}
}
