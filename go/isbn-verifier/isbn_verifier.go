package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(n string) bool {
	isbn := strings.ReplaceAll(n, "-", "")
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, c := range isbn {
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
