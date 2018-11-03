package main

import (
	"fmt"
)

func main() {
	student := make([]string, 10)
	students := make([][]string, 10)

	student[0] = "Rafael"
	student = append(student, "Alejandro")

	fmt.Println(student)
	fmt.Println(students)
}
