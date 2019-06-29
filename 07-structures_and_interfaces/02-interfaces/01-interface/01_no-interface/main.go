package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func main() {
	s1 := square{10}
	s2 := square{side: 12}

	fmt.Println("Area s1:", s1.area())
	fmt.Println("Area s2:", s2.area())
}
