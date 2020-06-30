// Package clock implements a solution for the clock exercism problem
package clock

import "fmt"

// Time represents the time with hours and minutes
type Time struct {
	hours, minutes int
}

func (t *Time) normalize() {
	t.hours = (t.hours + t.minutes/60) % 24
	t.minutes = t.minutes % 60

	if t.minutes < 0 {
		t.minutes += 60
		t.hours--
	}

	for t.hours < 0 {
		t.hours += 24
	}
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.hours, t.minutes)
}

// Add gives the time after m number of minutes
func (t Time) Add(m int) Time {
	t.minutes += m
	t.normalize()
	return t
}

// Subtract gives the time before m number of minutes
func (t Time) Subtract(m int) Time {
	return t.Add(-1 * m)
}

// New generates the time
func New(h, m int) Time {
	t := Time{
		hours:   h,
		minutes: m,
	}
	t.normalize()
	return t
}
