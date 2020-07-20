// Package letter implements a solution for the paralllel-letter-frequency exercism problem
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency concurrently counts the frequency of each rune in a given list of strings
func ConcurrentFrequency(l []string) FreqMap {
	result := FreqMap{}
	f := make(chan FreqMap, len(l))
	for _, s := range l {
		go cFrequency(s, f)
	}
	for range l {
		for r, c := range <-f {
			result[r] += c
		}
	}
	return result
}

func cFrequency(s string, c chan<- FreqMap) {
	c <- Frequency(s)
}

// Frequency counts the frequency of each rune in a given string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
