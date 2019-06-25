package main

import "fmt"

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

func main() {
	r := rect{width: 10, height: 10}
	fmt.Printf("%+v\n", r)
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())
	ptr := &rect{width: 10, height: 5}
	fmt.Println("Area: ", ptr.area())
	fmt.Println("Perimeter: ", ptr.perimeter())
}
