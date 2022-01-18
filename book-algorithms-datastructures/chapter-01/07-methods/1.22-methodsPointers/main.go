package main

import "fmt"

type myInt int

func (data myInt) increment1() {
	data = data + 1
}

func (data *myInt) increment2() {
	*data = *data + 1
}

func main() {
	var data myInt = 1
	fmt.Println("value before increment1() call :", data)
	data.increment1()
	fmt.Println("value after increment1() call :", data)
	fmt.Println("value before increment2() call :", data)
	data.increment2()
	fmt.Println("value after increment2() call :", data)
}
