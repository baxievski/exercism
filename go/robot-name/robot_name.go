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
	allCombinations := 26 * 26 * 10 * 10 * 10
	if len(usedNames) == allCombinations {
		return "", errors.New("all combinations for robot names have been exhausted")
	}
	if r.name != "" {
		return r.name, nil
	}
	r.name = randName()
	for usedNames[r.name] {
		r.name = randName()
	}
	usedNames[r.name] = true

	return r.name, nil
}

func randName() string {
	return string('A'+rand.Intn(26)) +
		string('A'+rand.Intn(26)) +
		string('0'+rand.Intn(10)) +
		string('0'+rand.Intn(10)) +
		string('0'+rand.Intn(10))
}
