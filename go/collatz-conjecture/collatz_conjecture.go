// Package collatzconjecture implements a solution for the collatz-conjecture exercism problem
package collatzconjecture

// package main

import (
	"fmt"
)

// CollatzConjecture returns the hamming distance between the two strings passed as arguments
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, fmt.Errorf("defined only for n > 0, not %v", n)
	}
	return collatzConjecture(n, 0), nil
}

func collatzConjecture(n int, step int) int {
	if n == 1 {
		return step
	}
	if n%2 == 0 {
		return collatzConjecture(n/2, step+1)
	}
	return collatzConjecture(3*n+1, step+1)
}

// func main() {
// 	fmt.Print(CollatzConjecture(120))
// 	fmt.Print(CollatzConjecture(1000000))
// }
