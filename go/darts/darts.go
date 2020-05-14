// Package darts implements a solution for the exercism darts challenge
package darts

import "math"

// Score returns the points earned in a toss of a darts game, given the coordinates of the hit
func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	if distance > 10 {
		return 0
	}
	if distance > 5 {
		return 1
	}
	if distance > 1 {
		return 5
	}
	return 10
}
