package main

import "fmt"

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Println(len(sf))
	fmt.Printf("%T \n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}
	fmt.Println(total)
	return total / float64(len(sf))
}

func main() {
	fmt.Println(average(43, 56, 87, 12, 45, 57))

	m := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(m)
}
