// Package grains implements a solution for the grains exercism problem
package grains

import (
	"fmt"
)

// Square gives the number of grains on the n-th square on a chess board
func Square(n int) (uint64, error) {
	if !(n >= 1 && n <= 64) {
		return 0, fmt.Errorf("defined only for 1 <= n <= 64, not %v", n)
	}
	return 1 << (n - 1), nil

}

// Total gives the number of grains on the whole chess board
func Total() uint64 {
	return 1<<64 - 1
}
