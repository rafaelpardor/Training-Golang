package main

import "fmt"

func incrementPassByPointer(ptr *int) {
	(*ptr)++
}

func main() {
	i := 10
	fmt.Println("Value of i before increment is : ", i)

	incrementPassByPointer(&i)
	fmt.Println("Value of i after increment is : ", i)
}
