package main

import "fmt"

func main() {
	maping := make(map[string]int)
	maping["Apple"] = 40
	maping["Banana"] = 30
	maping["Mango"] = 50

	for key, val := range maping {
		fmt.Printf("[ %v -> %v ]\n", key, val)
	}

	fmt.Println("Apple price:", maping["Apple"])

	delete(maping, "Apple")
	value, ok := maping["Apple"]
	fmt.Println("Apple price:", value, "Present:", ok)
	value2, ok2 := maping["Banana"]
	fmt.Println("Banana price:", value2, "Present:", ok2)
}
