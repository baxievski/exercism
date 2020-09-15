// Package twelve implements a solution for the twelve-days exercism problem
package twelve

import (
	"fmt"
	"strings"
)

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var presents = []string{
	"twelve Drummers Drumming, ",
	"eleven Pipers Piping, ",
	"ten Lords-a-Leaping, ",
	"nine Ladies Dancing, ",
	"eight Maids-a-Milking, ",
	"seven Swans-a-Swimming, ",
	"six Geese-a-Laying, ",
	"five Gold Rings, ",
	"four Calling Birds, ",
	"three French Hens, ",
	"two Turtle Doves, ",
	"and a Partridge in a Pear Tree.",
}

// Verse gives the verse from `The Twelve Days of Christmas` corresponding to the nth day
func Verse(n int) string {
	verse := "On the %v day of Christmas my true love gave to me: %v"
	if n == 1 {
		return fmt.Sprintf(verse, days[n-1], "a Partridge in a Pear Tree.")
	}

	return fmt.Sprintf(verse, days[n-1], strings.Join(presents[12-n:], ""))
}

// Song gives the `The Twelve Days of Christmas` song
func Song() string {
	s := make([]string, 12)
	for d := 1; d <= 12; d++ {
		s[d-1] = Verse(d)
	}
	return strings.Join(s, "\n")
}
