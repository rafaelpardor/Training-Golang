package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{"Walther", "White", 20}
	p2 := person{
		firstName: "Rafael",
		lastName:  "Pardo",
		age:       20,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("Hola, soy %s %s, tengo %d años\n", p2.firstName, p2.lastName, p2.age)
}
