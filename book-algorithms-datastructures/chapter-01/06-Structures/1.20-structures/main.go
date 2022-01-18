package main

import "fmt"

type student struct {
	rollNo int
	name   string
}

func main() {
	stud := student{1, "Johnny"}
	fmt.Println(stud)
	fmt.Println("Student name ::", stud.name)
	pstud := &stud
	fmt.Println("Student pointer name ::", pstud.name)
	fmt.Println(student{rollNo: 2, name: "Ann"})
	fmt.Println(student{name: "Gopher", rollNo: 2})
	fmt.Println(student{name: "Alice"})
}
