package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)
	fmt.Println(z.area())
}

func main() {
	s1 := square{side: 10}
	s2 := square{side: 20}

	info(s1)
	info(s2)
}
