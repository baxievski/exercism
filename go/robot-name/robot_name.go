// Package robotname implements a solution for the exercism robot-name challenge
package robotname

import (
	"errors"
	"math/rand"
	"time"
)

var usedNames = map[string]bool{}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Robot represents a robot with a name
type Robot struct {
	name string
}

// Reset resets the robot name
func (r *Robot) Reset() {
	r.name = ""
}

// Name generates a random name for the robot
func (r *Robot) Name() (string, error) {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	allCombinations := len(upper) * len(upper) * len(digits) * len(digits) * len(digits)

	if len(usedNames) == allCombinations {
		return "", errors.New("all combinations for robot names have been exhausted")
	}
	if r.name != "" {
		return r.name, nil
	}

	r.name = randString(upper, 2) + randString(digits, 3)
	for usedNames[r.name] {
		r.name = randString(upper, 2) + randString(digits, 3)
	}
	usedNames[r.name] = true

	return r.name, nil
}

func randString(choices string, length int) string {
	b := make([]rune, length)
	c := []rune(choices)
	for i := range b {
		b[i] = c[rand.Intn(len(c))]
	}
	return string(b)
}
