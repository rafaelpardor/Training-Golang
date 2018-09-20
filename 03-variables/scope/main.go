package main

import "fmt"

const globalVariable = "!!!-.-.-.-.-.-!!!I'm a global variable!!!-.-.-.-.-.-!!!"

func main() {
	fmt.Println("-.- Scopes -.-")
	fmt.Print("\n")

	fmt.Println(globalVariable) // Here

	globalVariable := 3
	fmt.Println("Diferent variable from the global variable.", globalVariable)

	varScope()
}

func varScope() {
	fmt.Print("\n")

	var x string

	x = "Python is good"
	fmt.Println(x)
	x = x + " but " + "Go is better"
	fmt.Println(x)

	fmt.Println(globalVariable) //
}
