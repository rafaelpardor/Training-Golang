// Function that returns a channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d\n", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	chn1 := boring("channel 1")
	chn2 := boring("channel 2")
	c := fanIn(chn1, chn2)
	for i := 0; i < 20; i++ {
		fmt.Printf(<-c)
	}
	fmt.Println("You are boring, I'm leaving")
}
