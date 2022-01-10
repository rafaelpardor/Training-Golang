package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person

	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", p1)
}
