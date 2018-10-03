package main

import (
	"fmt"

	"github.com/rafaelpardor/golang/01-basic/04-package/package1"
	// We can rename a package
	alias "github.com/rafaelpardor/golang/01-basic/04-package/package2"
)

func main() {
	x := "Rafael"
	y := "leafaR"
	foo := "!oG ,olleH"

	fmt.Println(stringutil.Exported)
	fmt.Println(alias.ExportedVariable)
	fmt.Println("String reversed:", stringutil.Reverse(x))
	fmt.Println("String reversed:", stringutil.Reverse(y))
	fmt.Println("String reversed:", stringutil.Reverse(foo))
}
