package stringutil

// A function that call another function from other file.Reverse
// This function will be exported.
func Reverse(s string) string {
	return reverseTwo(s)
}
