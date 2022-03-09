package main

import "fmt"

type Message struct {
	str  string
	wait chan bool
}

func main() {
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)

		msg1.wait <- true
	}
}
