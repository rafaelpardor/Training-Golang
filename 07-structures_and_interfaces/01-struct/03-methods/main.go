package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) printFullName() string {
	return fmt.Sprintf("Hola soy %s %s, y tego %d", p.firstName, p.lastName, p.age)
}

func main() {
	p1 := person{"Walther", "White", 20}
	p2 := person{
		firstName: "Rafael",
		lastName:  "Pardo",
		age:       20,
	}

	fmt.Println(p1.printFullName())
	fmt.Println(p2.printFullName())
}
