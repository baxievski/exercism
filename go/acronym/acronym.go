// Package acronym implements a solution for the exercism acronum challenge
package acronym

import (
	"unicode"
)

// Abbreviate retturns abbreviation of the passed string
func Abbreviate(s string) string {
	abbrivation := []rune{}
	prevRuneIsLetter := false
	sr := []rune(s)
	for _, char := range sr {
		if char == '\'' {
			continue
		}
		if !unicode.IsLetter(char) {
			prevRuneIsLetter = false
			continue
		}
		if !prevRuneIsLetter {
			abbrivation = append(abbrivation, unicode.ToUpper(char))
			prevRuneIsLetter = true
		}
	}
	return string(abbrivation)
}
