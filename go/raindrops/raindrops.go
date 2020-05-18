// Package raindrops implements a solution for the exercism raindrops chalenge
package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns the "raindrop" string corresponding to the factors of the integer arg
func Convert(num int) string {
	result := []string{}
	if num%3 == 0 {
		result = append(result, "Pling")
	}
	if num%5 == 0 {
		result = append(result, "Plang")
	}
	if num%7 == 0 {
		result = append(result, "Plong")
	}
	if len(result) == 0 {
		result = append(result, strconv.Itoa(num))
	}
	return strings.Join(result, "")

}
