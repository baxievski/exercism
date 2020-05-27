// Package isogram implements a solution for the exercism isogram challenge
package isogram

import "unicode"

// IsIsogram returns true if a word is an isogram
func IsIsogram(input string) bool {
	// ar, br := []rune(a), []rune(b)
	ir := []rune(input)
	seenChars := []rune{}
	for _, char := range ir {
		if !unicode.IsLetter(char) {
			continue
		}
		lowerChar := unicode.ToLower(char)
		if elemInSlice(lowerChar, seenChars) {
			return false
		}
		seenChars = append(seenChars, lowerChar)
	}
	return true
}

func elemInSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
