// Package dna implements a solution for the nucleotide-count exercism problem
package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
func (dna DNA) Counts() (Histogram, error) {
	histogram := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range dna {
		if _, ok := histogram[nucleotide]; ok {
			histogram[nucleotide]++
			continue
		}
		return Histogram{}, fmt.Errorf("no nucleotide %v in DNA", nucleotide)
	}

	return histogram, nil
}
