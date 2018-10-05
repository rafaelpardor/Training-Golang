package main

import "fmt"

// Now we combined all the types of create many variables at once.
// In this case, we create a way of declared many varaibles and with his own type.
func main() {
	var message = "Hello World!"
	var a, b, c = 1, false, "hi"
	var d, e, f string = `Hello`, "Rafael", "Pardo"
	var (
		g      = true
		h uint = 12
		i      = 'T'
	)

	fmt.Println(message, a, b, c, d, e, f, g, h, i)
}
