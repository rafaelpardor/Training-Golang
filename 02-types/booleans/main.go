package main

import "fmt"

func main() {
	fmt.Println("-.- Boolean operators -.-")
	getBooleans()
}

func getBooleans() {
	fmt.Print("\n")

	fmt.Println("true\t &&\t true.\t ", true)
	fmt.Println("true\t &&\t false.\t ", false)
	fmt.Println("false\t &&\t true.\t ", false)
	fmt.Println("false\t &&\t false.\t ", false)
	fmt.Print("\n")
	fmt.Println("true\t ||\t true.\t ", true)
	fmt.Println("true\t ||\t false.\t ", true)
	fmt.Println("false\t ||\t true.\t ", true)
	fmt.Println("false\t ||\t false.\t ", false)
	fmt.Print("\n")
	fmt.Println("!true\t", false)
	fmt.Println("!false\t", true)
	fmt.Print("\n")

	fmt.Println("2+2 == 4 && 53+2 == 55 ::", 2+2 == 4 && 53+2 == 55)
	fmt.Println("2+2 == 4 && 3*3 == 39", 2+2 == 4 && 3*3 == 39)
	fmt.Println("Platzi == Platzi || Rafael == Rafael ::", "Platzi" == "Platzi" || "Rafael" == "Rafael")
	fmt.Println("Platzi == Platzi || Go == Python ::", "Platzi" == "Platzi" || "Go" == "Python")
	fmt.Println("!true ::", !true)
	fmt.Println("!false ::", !false)
}
