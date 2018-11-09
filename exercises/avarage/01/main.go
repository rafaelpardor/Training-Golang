package main

import "fmt"

func main() {
	/*We create an array of length 5 to hold our
	test scores, then we fill up each element with a grade.*/
	var x [5]float64
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	var total float64
	/*We set up a for loop to compute the total score*/
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	/*we divide the total score by the number
	of elements to find the average*/
	fmt.Println(total / 5)
}
