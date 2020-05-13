// Package hamming implements a solution for the exercism hamming challenge
package hamming

import (
	"errors"
)

// Distance returns the hamming distance between the two strings passed as arguments
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("strings must have equal length")
	}
	result := 0
	for pos, char := range ar {
		if char != br[pos] {
			result++
		}
	}
	return result, nil
}
