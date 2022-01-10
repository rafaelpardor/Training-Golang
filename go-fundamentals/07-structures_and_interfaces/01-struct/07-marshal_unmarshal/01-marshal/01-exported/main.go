package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First          string
	Last           string
	Age            int
	intNotExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	bs, err := json.Marshal(p1)
	if err != nil {
		panic("AAAAAAAAAH")
	}

	fmt.Println("[]byte:", bs)
	fmt.Printf("%T es un alias para []byte \n", bs)
	fmt.Println(string(bs))
}
