package main

import (
	"fmt"
)

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	for i := 1; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizzBuzzOp() {
	for i := 1; i <= 100; i++ {
		result := ""
		if i%3 == 0 {
			result += "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}
		if result != "" {
			fmt.Println(result)
			continue
		}
		fmt.Println(i)
	}
}
