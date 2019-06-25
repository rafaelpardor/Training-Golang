package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)

	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	myGreeting["Rafael"] = "Buenos dias"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}
