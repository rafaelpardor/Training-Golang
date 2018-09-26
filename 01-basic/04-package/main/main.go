package main

import (
	"fmt"

	"github.com/rafaelpardor/golang/01-basic/04-package/package1"
	alias "github.com/rafaelpardor/golang/01-basic/04-package/package2"
)

func main() {
	x := "Rafael"
	y := "leafaR"
	foo := "!oG ,olleH"

	fmt.Println(alias.ExportedVariable)
	fmt.Println(stringutil.Exported)
	fmt.Println(stringutil.Reverse(x))
	fmt.Println(stringutil.Reverse(y))
	fmt.Println(stringutil.Reverse(foo))
}
