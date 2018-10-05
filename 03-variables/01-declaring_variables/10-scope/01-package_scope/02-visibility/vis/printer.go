package vis

import "fmt"

// The varialbes that start with a lowercase can be exported in the same env.
func PrintVar() {
	fmt.Println(Variable)
	fmt.Println(MyName)
	printVars()
}

// Print variables that is in another file
func printVars() {
	fmt.Println(yourName)
	fmt.Println(numbers)
	hola()
}
