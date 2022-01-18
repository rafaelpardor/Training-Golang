package main

import (
	"fmt"
	"math"
)

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	r1 := rect{width: 10, height: 10}
	c1 := circle{radius: 10}

	fmt.Printf("Rect: %+v\nCircle: %+v\n", r1, c1)
	fmt.Println("Reactangle total area: ", totalArea(r1))
	fmt.Println("Circle total area: ", totalArea(c1))
	fmt.Println("Total Perimeter: ", totalPerimeter(r1))
	fmt.Println("Total Perimeter: ", totalPerimeter(c1))
}
