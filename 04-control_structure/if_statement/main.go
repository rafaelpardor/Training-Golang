package main

import "fmt"

func main() {
	fmt.Println("-.- If Statement -.-")
	fmt.Println(more(1, 5))
	fmt.Println(max(10, 2))
	ifStatement()
	forIfStatement()
	fmt.Println(maxAreaCheck(1, 3, 3))
}

// Basic if
func more(x, y int) bool {
	fmt.Print("\n")

	if x > y {
		return true
	}
	return false

}

// If statement with else
func max(x, y int) int {
	fmt.Print("\n")

	var max int

	if x > y {
		max = x
	} else {
		max = y
	}
	return max
}

func ifStatement() {
	fmt.Print("\n")

	i := 11
	if i%2 == 0 {
		fmt.Println(i, "even")
	} else {
		fmt.Println(i, "odd")
	}
}

func forIfStatement() {
	fmt.Print("\n")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

// If with precondition.
func maxAreaCheck(length, width, limit int) string {
	fmt.Print("\n")

	if area := length * width; area > limit {
		return "The area is bigger than the limit"
	} else if area == limit {
		return "The area and the limit are equal."
	}
	return "The area is smaller than the limit"
}
