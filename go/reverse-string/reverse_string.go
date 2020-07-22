// Package reverse implements a solution for the exercism reverse-string challenge
package reverse

// Reverse reverses a string
func Reverse(s string) string {
	sr := []rune(s)
	length := len(sr)
	reversed := make([]rune, length)
	for i, v := range sr {
		reversed[length-1-i] = v
	}
	return string(reversed)
}
