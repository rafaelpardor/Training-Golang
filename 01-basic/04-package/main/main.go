package main

import (
	"fmt"

	alias "github.com/rafaelpardor/golang/01-basic/04-package/icomefromalaska"
	"github.com/rafaelpardor/golang/01-basic/04-package/stringutil"
)

func main() {
	x := "Rafael"
	y := "leafaR"
	foo := "!oG ,olleH"

	fmt.Println(alias.BearName)
	fmt.Println(stringutil.MyName)
	fmt.Println(stringutil.Reverse(x))
	fmt.Println(stringutil.Reverse(y))
	fmt.Println(stringutil.Reverse(foo))
}
