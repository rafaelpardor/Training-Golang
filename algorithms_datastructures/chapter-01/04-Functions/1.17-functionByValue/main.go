package main

import "fmt"

func incrementPassByValue(x int) {
	x++
}

func main() {
	i := 10
	fmt.Println("Value of i before increment is : ", i)

	incrementPassByValue(i)
	fmt.Println("Value of i after increment is : ", i)
}
