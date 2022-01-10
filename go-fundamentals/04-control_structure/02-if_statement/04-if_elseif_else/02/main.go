package main

func main() {
	maxAreaCheck(1, 2, 3)
	maxAreaCheck(2, 2, 4)
}

func maxAreaCheck(length, width, limit int) string {
	if area := length * width; area > limit {
		return "The area is bigger than the limit"
	} else if area == limit {
		return "The area and the limit are equal, it must be a square."
	}
	return "The area is smaller than the limit"
}
