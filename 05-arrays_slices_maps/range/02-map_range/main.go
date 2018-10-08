package main

import "fmt"

func main() {
	kvs := map[int]string{
		1: "apple",
		2: "banana",
	}
	for k, v := range kvs {
		fmt.Println(k, " -> ", v)
	}
}
