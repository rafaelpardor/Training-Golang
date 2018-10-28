package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",  //0
		"Bonjour!",       //1
		"dias!",          //2
		"Bongiorno!",     //3
		"Ohayo!",         //4
		"Selamat pagi!",  //5
		"Gutten morgen!", //6
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])
}
