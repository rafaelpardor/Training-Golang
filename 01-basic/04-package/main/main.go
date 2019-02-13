package main

import (
	"fmt"

	stringutil "github.com/rafaelpardor/golang/01-basic/04-package/package1"
	alias "github.com/rafaelpardor/golang/01-basic/04-package/package2"
)

func main() {
	x := "Rafael"
	y := "leafaR"
	foo := "!oG ,olleH"

	importing()
	fmt.Println(Hola) // variables.go
	fmt.Println(holi) // variables.go

	fmt.Println(alias.ExportedVariable)

	fmt.Println(stringutil.Exported)
	fmt.Println("String reversed:", stringutil.Reverse(x))
	fmt.Println("String reversed:", stringutil.Reverse(y))
	fmt.Println("String reversed:", stringutil.Reverse(foo))
}
