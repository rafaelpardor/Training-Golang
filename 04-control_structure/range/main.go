package main

import "fmt"

func main() {
	rangeStatement()
}

func rangeStatement() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 1}
	fmt.Printf("%T  \n", numbers)
	sum := 0
	for index, val := range numbers {
		sum += val
		fmt.Print("[", index, ",", val, "] ")

	}
	fmt.Println("\nSum is :: ", sum)

	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Println(k, " -> ", v)
	}

	str := "Hello, World!"
	for i, c := range str {
		fmt.Println("[", i, " ,", string(c), "] ")
	}
}
