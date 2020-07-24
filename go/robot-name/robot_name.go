// Package robotname implements a solution for the exercism robot-name challenge
package robotname

import (
	"errors"
	"math/rand"
	"time"
)

var usedNames = map[string]bool{}
var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

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
	if len(usedNames) == 26*26*10*10*10 {
		return "", errors.New("all robot names have been previously used")
	}
	if r.name != "" {
		return r.name, nil
	}

	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"

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
		b[i] = c[seed.Intn(len(c))]
	}
	return string(b)
}
