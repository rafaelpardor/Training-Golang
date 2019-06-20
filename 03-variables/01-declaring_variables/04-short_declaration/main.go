package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := "M"
	h, i, j, k := 'M', true, 12, "abcd"

	fmt.Println("Type \t||\t value")
	fmt.Println("---- \t\t -----")
	fmt.Printf("%T \t||\t %v \n", a, a)          // Type and value of variale a.
	fmt.Printf("%T \t||\t %v \n", b, b)          // Type and value of variale b.
	fmt.Printf("%T ||\t %v \n", c, c)            // Type and value of variale c.
	fmt.Printf("%T \t||\t %v \n", d, d)          // Type and value of variale d.
	fmt.Printf("%T \t||\t %v \n", e, e)          // Type and value of variale e.
	fmt.Printf("%T \t||\t %v \n", f, f)          // Type and value of variale f.
	fmt.Printf("%T \t||\t %v,%d \n", g, g, g[0]) // Type and value of variale g.
	fmt.Printf("%T \t||\t %v \n", h, h)          // Type and value of variale h.
	fmt.Printf("%T \t||\t %v \n", i, i)          // Type and value of variale i.
	fmt.Printf("%T \t||\t %v \n", j, j)          // Type and value of variale j.
	fmt.Printf("%T \t||\t %q \n", k, k)          // Type and value of variale k.
}
