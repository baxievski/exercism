// Package collatzconjecture implements a solution for the collatz-conjecture exercism problem
package collatzconjecture

import (
	"fmt"
)

// CollatzConjecture returns the number of steps for the 3x+1 problem
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, fmt.Errorf("defined only for n > 0, not %v", n)
	}
	steps := 0
	number := n
	for number > 1 {
		steps++
		if number%2 == 0 {
			number = number / 2
			continue
		}
		number = 3*number + 1
	}
	return steps, nil
}
