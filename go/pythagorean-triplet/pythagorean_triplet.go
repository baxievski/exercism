package pythagorean

import "math"

// Triplet represents sides of a triangle
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min, max
func Range(min, max int) []Triplet {
	result := []Triplet{}
	minSq := min * min
	maxSq := max * max
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			cSq := a*a + b*b
			if cSq < minSq {
				continue
			}
			if cSq > maxSq {
				continue
			}
			if !isPerfectSquare(cSq) {
				continue
			}
			c := int(math.Sqrt(float64(cSq)))
			result = append(result, Triplet{a, b, c})
		}
	}
	return result
}

// Sum returns a list of all Pythagorean triplets where the perimeter is equal to p
func Sum(p int) []Triplet {
	result := []Triplet{}
	triplets := Range(1, p/2+1)
	for _, t := range triplets {
		if t[0]+t[1]+t[2] != p {
			continue
		}
		result = append(result, t)
	}
	return result
}

func isPerfectSquare(n int) bool {
	if n < 2 {
		return true
	}

	x := n / 2
	for n < x*x {
		x = (x + n/x) / 2
	}

	return x*x == n
}
