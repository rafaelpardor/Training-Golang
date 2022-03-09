package main

import (
	"fmt"
	"time"
)

func incrementCount(n int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Go routine %v - %#v \n", n, i)
	}
}

func main() {
	fmt.Println("vim-go")
	//incrementCount(1)
	//incrementCount(2)
	go incrementCount(1)
	go incrementCount(2)
	go incrementCount(3)
	go incrementCount(4)
	go incrementCount(5)
	time.Sleep(7 * time.Second)
}
