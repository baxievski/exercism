// Package encode implements a solution for the exercism run-length-encoding challenge
package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode encodes a string with run length encoding
func RunLengthEncode(s string) string {
	runes := []rune(s)
	if len(runes) <= 1 {
		return s
	}
	var encoded strings.Builder
	prev := runes[0]
	reps := 1
	for i, char := range runes {
		if i == 0 {
			continue
		}

		if i < len(runes)-1 {
			if prev == char {
				reps++
				continue
			}
			if reps > 1 {
				encoded.WriteString(strconv.Itoa(reps))
			}
			encoded.WriteRune(prev)
			prev = char
			reps = 1
			continue
		}

		if prev == char {
			reps++
			encoded.WriteString(strconv.Itoa(reps))
			encoded.WriteRune(prev)
			continue
		}

		if reps > 1 {
			encoded.WriteString(strconv.Itoa(reps))
		}
		encoded.WriteRune(prev)
		encoded.WriteRune(char)
	}
	return encoded.String()
}

// RunLengthDecode decodes a run length encoded string
func RunLengthDecode(s string) string {
	runes := []rune(s)
	if len(runes) <= 1 {
		return s
	}
	var decoded strings.Builder
	ind := 0
	for i, char := range runes {
		if unicode.IsNumber(char) {
			continue
		}
		reps, err := strconv.Atoi(string(runes[ind:i]))
		if err != nil {
			decoded.WriteRune(char)
			ind = i + 1
			continue
		}
		decoded.WriteString(strings.Repeat(string(char), reps))
		ind = i + 1
	}
	return decoded.String()
}
