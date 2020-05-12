// Package accumulate implements a solution for the exercism accumulate chalenge
package accumulate

type function func(string) string

// Accumulate returns a slice with the result of applying f to each element from the input slice
func Accumulate(input []string, f function) []string {
	result := []string{}
	for _, i := range input {
		result = append(result, f(i))
	}
	return result
}
