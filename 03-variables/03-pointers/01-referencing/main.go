package main

import (
	"fmt"
)

func main() {
	var a int16 = 123

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Printf("Type %T\n", b)

}
