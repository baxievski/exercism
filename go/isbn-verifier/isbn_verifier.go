package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN validates a string that represents an isbn number
func IsValidISBN(isbn string) bool {
	n := strings.ReplaceAll(isbn, "-", "")
	if len(n) != 10 {
		return false
	}

	sum := 0
	for i, c := range n {
		d, err := strconv.Atoi(string(c))
		if err != nil {
			if i != 9 {
				return false
			}
			if c != 'X' {
				return false
			}
			d = 10
		}
		sum += d * (10 - i)
	}

	return sum%11 == 0
}
