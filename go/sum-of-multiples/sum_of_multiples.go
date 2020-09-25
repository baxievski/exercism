// Package summultiples implements a solution for the exercism sum-of-multiples chalenge
package summultiples

// SumMultiples gives the sum of all the numbers up to limit that are multiples of any of the divisors
func SumMultiples(limit int, divisors ...int) int {
	result := 0

	for n := 1; n < limit; n++ {
		for _, d := range divisors {
			if d == 0 {
				continue
			}
			if n%d == 0 {
				result += n
				break
			}
		}
	}

	return result
}
