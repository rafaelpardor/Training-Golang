package main

import "fmt"

func main() {

	if true && false {
		fmt.Println("This code will not run.")
	}

	fmt.Println("true\t &&\t true.\t ", true)
	fmt.Println("true\t &&\t false.\t ", false)
	fmt.Println("false\t &&\t true.\t ", false)
	fmt.Println("false\t &&\t false.\t ", false)
}
