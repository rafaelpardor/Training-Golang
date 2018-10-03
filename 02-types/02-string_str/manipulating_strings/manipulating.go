package main

import (
	"fmt"
)

func lengthOfStrings(str string) {
	fmt.Printf("The length of %q, is %d bits.\n", str, len(str))
}

func decimalString(str string) {
	fmt.Printf("The decimal of %q is %d.\n", str[2], str[2])
}

func replaceString(str string) {
	fmt.Printf("%s", str)
	fmt.Println(str, "/", str[0:2])
}
