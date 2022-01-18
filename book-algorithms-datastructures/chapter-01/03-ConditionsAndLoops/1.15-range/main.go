package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for index, val := range numbers {
		sum += val
		fmt.Printf("[%v,%v] ", index, val)
	}
	fmt.Println("\nSum of all numbers is ::", sum)

	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Println(k, "->", v)
	}

	str := "Hello, World!"
	for index, c := range str {
		fmt.Printf("[%v,%v]", string(c), index)
	}
}
