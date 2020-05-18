// Package raindrops implements a solution for the exercism raindrops chalenge
package raindrops

import (
	"strconv"
)

// Convert returns the "raindrop" string corresponding to the factors of the integer arg
func Convert(num int) string {
	result := ""
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		result += strconv.Itoa(num)
	}
	return result
}
