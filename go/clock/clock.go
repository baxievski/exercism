// Package clock implements a solution for the clock exercism problem
package clock

import "fmt"

// Time represents the time with hours and minutes
type Time struct {
	minutes int
}

// New generates the time
func New(hours, minutes int) Time {
	h := (hours + minutes/60) % 24
	m := minutes % 60

	for m < 0 {
		m += 60
		h--
	}

	for h < 0 {
		h += 24
	}

	return Time{
		minutes: h*60 + m,
	}
}

// Add gives the time after m number of minutes
func (t Time) Add(m int) Time {
	return New(0, t.minutes+m)
}

// Subtract gives the time before m number of minutes
func (t Time) Subtract(m int) Time {
	return New(0, t.minutes-m)
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.minutes/60, t.minutes%60)
}
