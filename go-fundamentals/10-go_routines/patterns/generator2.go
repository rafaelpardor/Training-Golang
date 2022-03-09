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
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	chn1 := boring("channel 1")
	chn2 := boring("channel 2")
	for i := 0; i < 5; i++ {
		fmt.Printf(<-chn1)
		fmt.Printf(<-chn2)
	}
	fmt.Println("You are boring, I'm leaving")
}
