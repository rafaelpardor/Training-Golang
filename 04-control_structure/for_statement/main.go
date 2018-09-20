package main

import "fmt"

func main() {
	fmt.Println("-.- For Statement -.-")
	forStatement()
	forSuma()
}

func forStatement() {
	fmt.Print("\n")

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Terminó el primer ciclo.")

	fmt.Print("\n")

	i := 11
	for i <= 20 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Terminó el segundo ciclo.")

	fmt.Print("\n")

	s := 20
	for {
		s++
		fmt.Println(s)
		if s == 30 {
			fmt.Println("Han terminado los ciclos.")
			break
		}
	}
}

func forSuma() {
	fmt.Print("\n")

	numbers := [...]int{1, 3, 4, 5, 6, 7, 8, 9, 10, 2}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(numbers)
	fmt.Println(sum)
}
