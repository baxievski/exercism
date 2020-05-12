// Package raindrops implements a solution for the exercism raindrops chalenge
package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns the "raindrop" string corresponding to the factors of the integer arg
func Convert(num int) string {
	result := []string{}
	f := factors(num)
	if elemInSlice(3, f) {
		result = append(result, "Pling")
	}
	if elemInSlice(5, f) {
		result = append(result, "Plang")
	}
	if elemInSlice(7, f) {
		result = append(result, "Plong")
	}
	if len(result) == 0 {
		result = append(result, strconv.Itoa(num))
	}
	return strings.Join(result, "")

}

func factors(num int) []int {
	result := []int{}
	counter := 2
	for counter*counter <= num {
		if num%counter != 0 {
			counter++
			continue
		}
		num = num / counter
		result = append(result, counter)
	}
	if num > 1 {
		result = append(result, num)
	}
	return result
}

func elemInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
