// Package triangle implements triangle exercism solution from the go track
package triangle

import (
	"math"
	"sort"
)

type Kind string

const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides takes the triangles sides and returns it's kind
func KindFromSides(a, b, c float64) Kind {
	if !isValidTriangle(a, b, c) {
		return NaT
	}
	if (a == b) && (b == c) {
		return Equ
	}
	if (a == b) || (b == c) || (a == c) {
		return Iso
	}
	return Sca
}

func isValidTriangle(a, b, c float64) bool {
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	for _, side := range sides {
		if (side <= 0) || (side == math.Inf(1)) {
			return false
		}
	}
	return sides[0]+sides[1] >= sides[2]
}
