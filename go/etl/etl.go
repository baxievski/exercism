// Package etl implements a solution for the exercism etl chalenge
package etl

import "strings"

// Transform flattens the map of lists into a single map
func Transform(input map[int][]string) map[string]int {
	result := map[string]int{}
	for k, l := range input {
		for _, v := range l {
			result[strings.ToLower(v)] = k
		}
	}
	return result
}
