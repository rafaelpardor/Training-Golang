package main

import "fmt"

func main() {
	fmt.Println("-.- Switch Statement -.-")
	switchStatement()
	isEven(2)
	fowSwitchStatemet()
	switchTest()
}

func switchStatement() {
	fmt.Print("\n")

	i := 10
	fmt.Println(i)
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3, 4, 5:
		fmt.Println("Three, Four, Five")
	case 10:
		fmt.Println("Number 10")
	default:
		fmt.Printf("%d is not in the in cases.\n", i)
	}
}

func isEven(value int) {
	fmt.Print("\n")

	switch {
	case value%2 == 0:
		fmt.Println(value, "Is even")
	default:
		fmt.Println(value, "Is odd")
	}
}

func fowSwitchStatemet() {
	fmt.Print("\n")

	for i := 0; i < 11; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "Es par")
		}
	}
}

func switchTest() {
	fmt.Print("\n")

	var number int

	fmt.Print("Ingresa un numero del 1 al 10: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("El numero ingresado es uno")
	default:
		fmt.Println("El numero ingresado no es 1")
	}

	switch {
	case number%2 == 0:
		fmt.Println("Además es par")
	default:
		fmt.Println("Además es impar")
	}
}
