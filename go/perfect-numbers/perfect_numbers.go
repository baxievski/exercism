package perfect

import (
	"fmt"
)

// Classification holds the classification values for natural numbers
type Classification int

const (
	// ClassificationAbundant represents abundant numbers
	ClassificationAbundant Classification = 1

	// ClassificationDeficient represents deficient numbers
	ClassificationDeficient Classification = 2

	// ClassificationPerfect represents perfect numbers
	ClassificationPerfect Classification = 3
)

// ErrOnlyPositive is an error...
var ErrOnlyPositive error = fmt.Errorf("cannot classify non-positive numbers")

// Classify tells if a number is perfect, abundant, or deficient
// based on Nicomachus' classification, errors for non-positive numbers
func Classify(n int64) (Classification, error) {
	var c Classification
	if n <= 0 {
		return c, ErrOnlyPositive
	}
	f := factors(n)
	s := sum(f)
	if s < 2*n {
		return ClassificationDeficient, nil
	}
	if s > 2*n {
		return ClassificationAbundant, nil
	}
	return ClassificationPerfect, nil
}

func factors(n int64) []int64 {
	f := map[int64]bool{}
	for i := int64(1); i*i < n; i++ {
		if n%i == 0 {
			f[i] = true
			f[n/i] = true
		}
	}
	r := make([]int64, 0, len(f))
	for k := range f {
		r = append(r, k)
	}
	return r
}

func sum(n []int64) int64 {
	s := int64(0)
	for _, i := range n {
		s += i
	}
	return s
}
