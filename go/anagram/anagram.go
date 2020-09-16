package anagram

import (
	"sort"
	"strings"
)

// Detect checks which of the strings in the list of candidates are anagrams of the input word
func Detect(s string, candidates []string) []string {
	result := []string{}
	for _, c := range candidates {
		if IsAnagram(s, c) {
			result = append(result, c)
		}
	}
	return result
}

// IsAnagram tells if the two strings are anagrams of each other
func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	lA := strings.ToLower(a)
	lB := strings.ToLower(b)
	if lA == lB {
		return false
	}
	sA := []rune(lA)
	sort.Slice(sA, func(i, j int) bool { return sA[i] < sA[j] })
	sB := []rune(lB)
	sort.Slice(sB, func(i, j int) bool { return sB[i] < sB[j] })
	for i, v := range sA {
		if v != sB[i] {
			return false
		}
	}
	return true
}
