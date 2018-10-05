package main

import (
	"fmt"

	"github.com/rafaelpardor/golang/03-variables/01-declaring_variables/10-scope/01-package_scope/02-visibility/vis"
)

func main() {
	fmt.Println(vis.Variable)
	fmt.Println(vis.MyName)
	fmt.Println("-Exporting functions-")
	vis.PrintVar()
}
