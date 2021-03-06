package twobucket

import "fmt"

type bucket struct {
	size  int
	level int
}

func Solve(b1Size, b2Size, goal int, startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error) {
	if b1Size < goal && b2Size < goal {
		return "", 0, 0, fmt.Errorf("no solution")
	}
	if goal%gcd(b1Size, b2Size) != 0 {
		return "", 0, 0, fmt.Errorf("no solution")
	}
	return "", 0, 0, nil
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
