package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-.- Strings manipulation -.-")
	getName()
	getStrings()
	strings2()
	japanese := getUnicode()
	fmt.Println(japanese)
}

func getName() {
	fmt.Print("\n")

	var name string

	fmt.Print("Please write your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello your name is, %s, welcome to Go!.\n", name)
}

func getStrings() {
	fmt.Print("\n")

	x := "Golang."

	fmt.Printf("Used to find the number of characters in the strings :: '%v' is %d. \n", x, len(x))
	fmt.Printf("Acceso a un bit/caracter en particular en un string :: %d. \n ", x[0])
	fmt.Printf("Acceso a un caracter en particular en un string :: %s. \n", string(x[0]))
	fmt.Println("Two strings are concatenated into a single string ::", "Hello,"+" world!")
	fmt.Println(x, x[0:2])
}

func strings2() {
	fmt.Print("\n")

	var text = "Hello world, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Golang", -1))
	fmt.Println(strings.Split(text, "H"))
}

func getUnicode() string {
	fmt.Print("\n")

	moshimoshi := "もしもし"
	return moshimoshi
}
