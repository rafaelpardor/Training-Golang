package main

import (
	"fmt"

	"github.com/rafaelpardor/golang/01-basic/04-package/package1"
	// We can put an alias ti the packages
	alias "github.com/rafaelpardor/golang/01-basic/04-package/package2"
)

func main() {
	x := "Rafael"
	y := "leafaR"
	foo := "!oG ,olleH"

	importing()
	fmt.Println(Hola) //Variables.go
	fmt.Println(holi) //Variables.go

	fmt.Println(alias.ExportedVariable)

	fmt.Println(stringutil.Exported)
	fmt.Println("String reversed:", stringutil.Reverse(x))
	fmt.Println("String reversed:", stringutil.Reverse(y))
	fmt.Println("String reversed:", stringutil.Reverse(foo))
}
