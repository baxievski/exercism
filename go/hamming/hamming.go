// Package hamming implements a solution for the exercism hamming chalenge
package hamming

import (
	"errors"
)

// Distance returns the hamming distance between the two strings passed as arguments
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings must have equal length")
	}
	result := 0
	for pos, char := range a {
		if char != []rune(b)[pos] {
			result++
		}
	}
	return result, nil
}
