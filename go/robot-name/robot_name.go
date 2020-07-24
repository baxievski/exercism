package robotname

// package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var usedNames = map[string]bool{}
var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func (r *Robot) Reset() {
}

func (r *Robot) Name() (string, error) {
	if len(usedNames) == 26*26*10*10*10 {
		return "", errors.New("text string")
	}
	if r.name != "" {
		return r.name, nil
	}
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	r.name = randString(2, upper) + randString(3, digits)
	for usedNames[r.name] {
		r.name = randString(2, upper) + randString(3, digits)
	}
	fmt.Println(r.name)
	usedNames[r.name] = true
	return r.name, nil
}

func randString(length int, choices string) string {
	b := make([]rune, length)
	c := []rune(choices)
	for i := range b {
		b[i] = c[seed.Intn(len(c))]
	}
	return string(b)
}
