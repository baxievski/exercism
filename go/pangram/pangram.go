// Package pangram implements a solution for the exercism pangram challenge
package pangram

import "unicode"

// IsPangram tells if the input string contains all ASCII characters
func IsPangram(i string) bool {
	seen := map[rune]bool{}
	for _, r := range []rune(i) {
		seen[unicode.ToLower(r)] = true
	}
	for r := 'a'; r <= 'z'; r++ {
		if _, ok := seen[r]; !ok {
			return false
		}
	}
	return true
}
