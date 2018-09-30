package main

func main() {
	max(1, 2)
}
func max(x, y int) int {
	var max int

	if x > y {
		max = x
	} else {
		max = y
	}
	return max
}
