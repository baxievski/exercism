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
	for i, rn := range reverse(input) {
		if !unicode.IsDigit(rn) {
			return 0, errors.New("input can only have digits and optional spaces")
		}
		digit := int(rn - '0')

		if i%2 == 0 {
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

func reverse(input []rune) []rune {
	output := make([]rune, len(input))
	copy(output, input)

	for i := len(output)/2 - 1; i >= 0; i-- {
		j := len(output) - 1 - i
		output[i], output[j] = output[j], output[i]
	}

	return output
}
