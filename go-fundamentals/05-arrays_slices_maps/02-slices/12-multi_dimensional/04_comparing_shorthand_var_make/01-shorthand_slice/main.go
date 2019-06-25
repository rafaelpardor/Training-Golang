package main

import (
	"fmt"
)

func main() {
	student := [1]string{}
	students := [][]string{}

	student[0] = "Todd"
	// student = append(student, "aa")

	fmt.Printf("%T", student)
	fmt.Println(students)
}
