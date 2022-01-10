package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Buenos dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}
	fmt.Printf("\n")
	for j := 0; j < len(greeting); j++ {
		fmt.Println(j, greeting[j])
	}
}
