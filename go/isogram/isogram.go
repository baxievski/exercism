// Package isogram implements a solution for the exercism isogram challenge
package isogram

import "unicode"

// IsIsogram returns true if a word is an isogram
func IsIsogram(input string) bool {
	ir := []rune(input)
	seenChars := map[rune]bool{}
	for _, char := range ir {
		if !unicode.IsLetter(char) {
			continue
		}
		lowerChar := unicode.ToLower(char)
		if seenChars[lowerChar] {
			return false
		}
		seenChars[lowerChar] = true
	}
	return true
}
