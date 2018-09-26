package main

import "fmt"

func main() {

	if true || false {
		fmt.Println("This ran.")
	}
	if false || false {
		fmt.Println("This not.")
	}

	fmt.Println("true\t ||\t true.\t ", true)
	fmt.Println("true\t ||\t false.\t ", true)
	fmt.Println("false\t ||\t true.\t ", true)
	fmt.Println("false\t ||\t false.\t ", false)
}
