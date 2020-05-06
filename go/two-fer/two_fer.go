// Package twofer implements the Exercism "two-fer" exercise.
package twofer

import "fmt"

// ShareWith returns a string in the form "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
