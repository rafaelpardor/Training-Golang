package main

import "fmt"

func main() {
	kvs := map[int]string{
		1: "apple",
		2: "banana",
		3: "pear",
	}

	for k, v := range kvs {
		fmt.Printf("%d -> %v\n", k, v)
	}
}
