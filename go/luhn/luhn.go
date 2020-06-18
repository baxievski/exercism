// Package luhn implements a solution for the exercism luhn challenge
package luhn

import (
	"errors"
	"strings"
	"unicode"
)

// Valid tells if the input string is a valid luhn number
func Valid(input string) bool {
	inputR := []rune(strings.ReplaceAll(input, " ", ""))

	if len(inputR) <= 1 {
		return false
	}

	l, err := luhn(inputR)

	if err != nil {
		return false
	}

	if l%10 != 0 {
		return false
	}

	return true
}

func luhn(input []rune) (int, error) {
	result := 0
	for i, rn := range input {
		if !unicode.IsDigit(rn) {
			return 0, errors.New("input can only have digits and optional spaces")
		}
		digit := int(rn - '0')

		if (len(input)+i-1)%2 == 0 {
			result += digit
			continue
		}

		digitDoubled := 2 * digit
		if digitDoubled > 9 {
			digitDoubled -= 9
		}
		result += digitDoubled
	}
	return result, nil
}
