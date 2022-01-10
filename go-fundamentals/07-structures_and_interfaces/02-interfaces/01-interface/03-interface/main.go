package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Printf("El area de la figura es: %+v\n", z)
	fmt.Println(z.area())
}

func main() {
	s1 := square{10}
	c1 := circle{5}

	info(s1)
	info(c1)
}
