package main

import "fmt"

func main() {
	/*
		The syntax of switch basic form
		switch <Initialization>; <condition> {
		case <value1>:
			<statements>
		case<value2>:
			<statements>
		We can have any number of case statements
		default: Optional
			<statements>
		}
	*/
	switch value := 0; value == 1 {
	case value%2 == 0:
		fmt.Println("valie is even")
	case value%2 == 1:
		fmt.Println("value is odd")
	default:
	}
}
