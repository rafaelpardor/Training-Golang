package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	greeting = append(greeting, "Suprabadham")

	fmt.Println(cap(greeting))
	fmt.Println(greeting[3])
}
