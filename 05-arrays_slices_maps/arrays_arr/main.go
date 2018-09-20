package main

import "fmt"

func main() {
	fmt.Println("-.- Arrays -.-")
	basicArray()
	avarageArray1()
	avarageArray2()
	avarageArray3()
}

func basicArray() {
	fmt.Print("\n")

	var x [5]int
	x[1] = 1
	x[4] = 105

	fmt.Println(x)
	fmt.Println(x[4])
}

func avarageArray1() {
	fmt.Print("\n")

	var x [10]float64
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5
	x[5] = 6
	x[6] = 7
	x[7] = 8
	x[8] = 9
	x[9] = 10
	var total float64

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total)
	fmt.Println(total / float64(len(x)))
}

func avarageArray2() {
	fmt.Print("\n")

	x := [12]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var total float64

	for _, value := range x {
		total += value
	}
	fmt.Println(total)
	fmt.Println(total / float64(len(x)))
}

func avarageArray3() {
	fmt.Print("\n")

	x := [3]float64{1, 1, 1}
	fmt.Println(x)
	var total float64

	for _, value := range x {
		total += value
	}
	fmt.Println(total)
	fmt.Println(total / float64(len(x)))

}

func getArray() {
	var arr1 [2]string
	arr1[0] = "Rafael"
	arr1[1] = "Freddy"
	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
