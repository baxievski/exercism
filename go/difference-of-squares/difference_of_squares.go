// Package diffsquares implements a solution for the exercism difference-of-squares challenge
package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return (2*n + 1) * n * (n + 1) / 6
}

// Difference returns the difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
