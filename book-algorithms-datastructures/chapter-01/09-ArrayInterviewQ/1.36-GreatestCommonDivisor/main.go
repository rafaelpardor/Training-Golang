package main

import "fmt"

func main() {
	fmt.Println(2, 4)
}

func gcd(m, n int) int {
	if m < n {
		return gcd(n, m)
	}

	if m%n == 0 {
		return n
	}
	return gcd(n, m%n)
}
