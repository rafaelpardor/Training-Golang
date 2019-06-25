package main

func permutation(data []int, i int, length int) {
	if length == i {
		return
	}
	for j := i; j < length; j++ {
		swap(data, i, j)
		permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func main() {
	var data [5]int
	for i := 0; i < 5; i++ {
		data[i] = i
	}
	permutation(data[:], 0, 5)
}
