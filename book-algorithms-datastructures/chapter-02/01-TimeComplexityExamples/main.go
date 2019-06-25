package main

import (
	"fmt"
	"time"
)

// Time-Complexity: O(n)
func bigO1(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m++
	}
	return m
}

//  TIme-Complexity: O(n^2)
func bigO2(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m++
		}
	}
	return m
}

// Time-Complexity: O(N+(N-1)+(N-2)+...) == O(N(N+1)/2) == O(n^2)
func bigO3(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			m++
		}
	}
	return m
}

// Each time problem space is divided into half. Time Complexity: O(log(n))
func bigO4(n int) int {
	m := 0
	i := 1
	for i < n {
		m++
		i = i * 2
	}
	return m
}

// Same as above each time problem space is divided into half. Time Complexity: O(log(n))
func bigO5(n int) int {
	m := 0
	i := n
	for i > 0 {
		m++
		i = i / 2
	}
	return m
}

// Outer loop will run for n number of iterations. In each iteration of the outer
// loop, inner loop will run for n iterations of their own. Final complexity will be
// n*n*n
func bigO6(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				m++
			}
		}
	}
	return m
}

func main() {
	startTime := time.Now()

	value := 10
	fmt.Printf("%v instructions for O(n) :: %v\n", value, bigO1(value))
	fmt.Printf("%v instructions for O(n^2) :: %v\n", value, bigO2(value))
	fmt.Printf("%v instructions for O(n^2) :: %v\n", value, bigO3(value))
	fmt.Printf("%v instructions for O(log(n)) :: %v\n", value, bigO4(value))
	fmt.Printf("%v instructions for O(log(n)) :: %v\n", value, bigO6(value))
	fmt.Printf("%v instructions for O(n^3) :: %v\n", value, bigO5(value))

	elapsed := time.Since(startTime)
	fmt.Printf("%s\n", elapsed)
}
