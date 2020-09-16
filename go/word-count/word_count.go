package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map of the number of occurances of each word
type Frequency map[string]int

// WordCount counts the occurances of each word in the input
func WordCount(i string) Frequency {
	result := Frequency{}
	r := regexp.MustCompile(`[^a-z0-9']+`)
	words := strings.Fields(r.ReplaceAllString(strings.ToLower(i), " "))
	for _, w := range words {
		result[strings.Trim(w, "'")]++
	}
	return result
}
