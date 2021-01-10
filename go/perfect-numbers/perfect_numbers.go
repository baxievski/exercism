package perfect

import (
	"fmt"
)

// Classification holds the classification values for natural numbers
type Classification string

const (
	// ClassificationAbundant represents abundant numbers
	ClassificationAbundant Classification = "abundant"

	// ClassificationDeficient represents deficient numbers
	ClassificationDeficient Classification = "deficient"

	// ClassificationPerfect represents perfect numbers
	ClassificationPerfect Classification = "perfect"
)

// ErrOnlyPositive is an error...
var ErrOnlyPositive error = fmt.Errorf("cannot classify non-positive numbers")

// Classify tells if a number is perfect, abundant, or deficient
// based on Nicomachus' classification, errors for non-positive numbers
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		var c Classification
		return c, ErrOnlyPositive
	}

	sum := aliquotSum(n)
	if sum < n {
		return ClassificationDeficient, nil
	}
	if sum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationPerfect, nil
}

func aliquotSum(n int64) int64 {
	factors := map[int64]bool{}

	for i := int64(1); i*i < n; i++ {
		if n%i != 0 {
			continue
		}
		factors[i] = true
		factors[n/i] = true
	}
	factors[n] = false

	sum := int64(0)
	for i, ok := range factors {
		if !ok {
			continue
		}
		sum += i
	}
	return sum
}
