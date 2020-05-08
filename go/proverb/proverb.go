// Package proverb implements a solution for the proverb exercism problem
package proverb

import "fmt"

// Proverb generates a proverb from the list of items passed as argument
func Proverb(rhyme []string) []string {
	result := []string{}
	regLine := "For want of a %v the %v was lost."
	endLine := "And all for the want of a %v."
	for i, w := range rhyme {
		if i+1 < len(rhyme) {
			result = append(result, fmt.Sprintf(regLine, w, rhyme[i+1]))
			continue
		}
		result = append(result, fmt.Sprintf(endLine, rhyme[0]))
	}
	return result
}
