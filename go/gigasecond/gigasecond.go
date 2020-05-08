// Package gigasecond implements a solution for the exercism gigasecond chalenge
package gigasecond

import (
	"time"
)

// AddGigasecond returns the time 10^9 seconds after the input time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
