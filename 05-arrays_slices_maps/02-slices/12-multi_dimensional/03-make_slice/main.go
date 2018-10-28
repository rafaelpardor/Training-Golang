package main

import (
	"fmt"
)

func main() {
	student := make([]string, 10)
	students := make([][]string, 10)

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println(students == nil)
}
