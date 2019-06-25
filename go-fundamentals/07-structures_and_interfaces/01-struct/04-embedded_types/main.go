package main

import (
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	age       int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) printFullName() string {
	return fmt.Sprintf("Hola soy %s %s, y tengo %d", p.FirstName, p.LastName, p.age)
}

func main() {
	p1 := doubleZero{
		person: person{
			FirstName: "James",
			LastName:  "Bond",
			age:       20,
		},
		LicenseToKill: true,
	}
	p2 := doubleZero{
		person: person{
			FirstName: "Miss",
			LastName:  "MoneyPenny",
			age:       19,
		},
		LicenseToKill: false,
	}
	p3 := doubleZero{
		person: person{
			"rafael",
			"Pardo",
			12,
		},
	}

	fmt.Println(p1.FirstName, p1.LastName, p1.age, p1.LicenseToKill)
	fmt.Println(p2.FirstName, p2.LastName, p2.age, p2.LicenseToKill)
	fmt.Printf("%v\n", p3.printFullName())
	fmt.Printf("%#v\n", p3)
	fmt.Printf("%+v\n", p3)
}
