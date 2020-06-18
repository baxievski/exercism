// Package luhn implements a solution for the exercism luhn challenge
package luhn

import (
	"strings"
	"unicode"
)

// Valid tells if the input string is a valid luhn number
func Valid(input string) bool {
	inputR := []rune(strings.ReplaceAll(input, " ", ""))

	if len(inputR) <= 1 {
		return false
	}

	luhn := 0
	for i, rn := range inputR {
		if !unicode.IsDigit(rn) {
			return false
		}

		d := int(rn - '0')
		if (len(inputR)+i-1)%2 == 0 {
			luhn += d
			continue
		}

		doubled := 2 * d
		if doubled > 9 {
			doubled -= 9
		}
		luhn += doubled
	}

	if luhn%10 != 0 {
		return false
	}

	return true
}
