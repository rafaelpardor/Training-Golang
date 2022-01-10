package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings package and string methods")

	fmt.Println(strings.ToTitle("hola"))
	fmt.Println(strings.Contains("Test", "t"))
	fmt.Println(strings.Count("aaa", "a"))
	fmt.Println(strings.HasPrefix("Rafael", "R"))
	fmt.Println(strings.Index("Rafael", "r"))
	fmt.Println(strings.Join([]string{"", "R", "a", "f", "a", "e", "l", ""}, "|"))
	fmt.Println(strings.Repeat("AP", 5))
	fmt.Println(strings.Replace("abcdefgghi", "g", "b", 2))
	fmt.Println(strings.Split("R-a-f-a-e-l", "-"))
}
