package main

import "fmt"

func main() {
	fmt.Println("In Go we can create a variable with a short declaration, using the Walrus operator, with this declaration, Go will infer the data type..")

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
	fmt.Printf("%T \t||\t %v \n", a, a)
	fmt.Printf("%T \t||\t %v \n", b, b)
	fmt.Printf("%T ||\t %v \n", c, c)
	fmt.Printf("%T \t||\t %v \n", d, d)
	fmt.Printf("%T \t||\t %v \n", e, e)
	fmt.Printf("%T \t||\t %v \n", f, f)
	fmt.Printf("%T \t||\t %v,%d \n", g, g, g[0])
	fmt.Printf("%T \t||\t %v \n", h, h)
	fmt.Printf("%T \t||\t %v \n", i, i)
	fmt.Printf("%T \t||\t %v \n", j, j)
	fmt.Printf("%T \t||\t %q \n", k, k)
}
