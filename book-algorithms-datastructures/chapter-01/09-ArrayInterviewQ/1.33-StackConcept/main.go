package main

import "fmt"

func function1() {
	fmt.Println("Func1 line 1")
	function2()
	fmt.Println("Func1 line 2")
}

func function2() {
	fmt.Println("Func2 line 1")
}

func main() {
	fmt.Println("Main line 1")
	function1()
	fmt.Println("Main line 2")
}
