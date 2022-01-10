package main

import "fmt"

var a = "stored in a"
var b, c string = "stored in b", "stored in c"
var d string

func main() {
	d = "stored in d"
	var e = 42
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := `o`

	fmt.Printf("a - %v\n.", a)
	fmt.Printf("a - %v\n.", b)
	fmt.Printf("a - %v\n.", c)
	fmt.Printf("a - %v\n.", d)
	fmt.Printf("a - %v\n.", e)
	fmt.Printf("a - %v\n.", f)
	fmt.Printf("a - %v\n.", g)
	fmt.Printf("a - %v\n.", h)
	fmt.Printf("a - %v\n.", i)
	fmt.Printf("a - %v\n.", j)
	fmt.Printf("a - %v\n.", k)
	fmt.Printf("a - %v\n.", l)
	fmt.Printf("a - %v\n.", m)
	fmt.Printf("a - %v\n.", n)
	fmt.Printf("a - %v\n.", o)
}
