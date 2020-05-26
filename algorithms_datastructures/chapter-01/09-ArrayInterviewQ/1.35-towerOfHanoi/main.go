package main

import "fmt"

func towersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are:")
	tohutil(num, "A", "C", "B")
}

func tohutil(num int, from, to, temp string) {
	if num < 1 {
		return
	}

	tohutil(num-1, from, temp, to)
	fmt.Println("Move disk", num, "from peg", from, "to peg", to)
	tohutil(num-1, temp, to, from)
}

func main() {
	towersOfHanoi(3)
}
