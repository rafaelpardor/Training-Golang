package vis

import "fmt"

func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
	printVars()
}

func printVars() {
	fmt.Println(Variable)
}
