// Package bob implements a solution for the exercism bob chalenge
package bob

import (
	"strings"
	"unicode"
)

// Hey returns the Bob's response to a remark
func Hey(remark string) string {
	remark = strings.Trim(remark, "\t\n\r ")
	if remark == "" {
		return "Fine. Be that way!"
	}
	if isUpper(remark) && endsWith(remark, "?") {
		return "Calm down, I know what I'm doing!"
	}
	if isUpper(remark) {
		return "Whoa, chill out!"
	}
	if endsWith(remark, "?") {
		return "Sure."
	}
	return "Whatever."
}

func isUpper(s string) bool {
	containsLetter := false
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		if !unicode.IsUpper(r) {
			return false
		}
		containsLetter = true
	}
	return containsLetter
}

func endsWith(s, c string) bool {
	return string(s[len(s)-1]) == c
}
