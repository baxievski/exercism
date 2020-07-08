// Package clock implements a solution for the clock exercism problem
package clock

import "fmt"

// Time represents the time with hours and minutes
type Time struct {
	hours, minutes int
}

// New generates the time
func New(hours, minutes int) Time {
	hours = (hours + minutes/60) % 24
	minutes = minutes % 60

	if minutes < 0 {
		minutes += 60
		hours--
	}

	for hours < 0 {
		hours += 24
	}

	return Time{
		hours:   hours,
		minutes: minutes,
	}
}

// Add gives the time after m number of minutes
func (t Time) Add(m int) Time {
	return New(t.hours, t.minutes+m)
}

// Subtract gives the time before m number of minutes
func (t Time) Subtract(m int) Time {
	return New(t.hours, t.minutes-m)
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.hours, t.minutes)
}
