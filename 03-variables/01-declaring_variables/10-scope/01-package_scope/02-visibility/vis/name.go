package vis

import (
	"fmt"
)

// The variables that start with uppercase will be exported
var Variable = "This varialbe will be exported"
var MyName = "Rafael"
var yourName = "golang lover"
var numbers int16 = 1986

func hola() {
	c := "Varialbe exported from a scope function"
	fmt.Println(c)
}
