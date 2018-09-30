package main

func main() {
	maxAreaCheck(1, 2, 3)
}

func maxAreaCheck(length, width, limit int) string {
	if area := length * width; area > limit {
		return "The area is bigger than the limit"
	} else if area == limit {
		return "The area and the limit are equal."
	}
	return "The area is smaller than the limit"
}
