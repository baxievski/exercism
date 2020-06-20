// Package luhn implements a solution for the exercism luhn challenge
package luhn

import (
	"strings"
	"unicode"
)

// Valid tells if the input string is a valid luhn number
func Valid(input string) bool {
	inp := strings.ReplaceAll(input, " ", "")

	if len(inp) <= 1 {
		return false
	}

	double := len(inp)%2 == 0
	luhn := 0
	for _, r := range inp {
		if !unicode.IsDigit(r) {
			return false
		}

		d := int(r - '0')
		if double {
			d = 2 * d
			if d > 9 {
				d -= 9
			}
		}

		luhn += d
		double = !double
	}

	return luhn%10 == 0
}
