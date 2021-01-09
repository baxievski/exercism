package lsproduct

import (
	"fmt"
)

// LargestSeriesProduct gives the largest product for contiguous string of digis of length n
func LargestSeriesProduct(digits string, n int) (int64, error) {
	if n < 0 {
		return 0, fmt.Errorf("span must be greater than zero")
	}
	if n > len(digits) {
		return 0, fmt.Errorf("span must be smaller than string length")
	}

	prod := int64(0)
	for i := 0; i <= len(digits)-n; i++ {
		p := int64(1)
		for _, d := range digits[i : i+n] {
			if d < '0' || d > '9' {
				return 0, fmt.Errorf("digits input must only contain digits")
			}
			p *= int64(d - '0')
		}
		if p > prod {
			prod = p
		}
	}

	return prod, nil
}
