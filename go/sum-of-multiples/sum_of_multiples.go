// Package summultiples implements a solution for the exercism sum-of-multiples chalenge
package summultiples

// SumMultiples gives the sum of all the numbers up to limit that are multiples of any of the divisors
func SumMultiples(limit int, divisors ...int) int {
	multiples := map[int]bool{}
	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for j := 1; j < limit; j++ {
			if j%d == 0 {
				multiples[j] = true
			}
		}
	}
	result := 0
	for m := range multiples {
		result += m
	}
	return result
}
