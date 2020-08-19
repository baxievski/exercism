// Package twelve implements a solution for the twelve-days exercism problem
package twelve

import (
	"fmt"
	"strings"
)

type dailyPresent struct {
	day     string
	present string
}

// Verse gives the verse from `The Twelve Days of Christmas` corresponding to the nth day
func Verse(n int) string {
	dailyPresents := map[int]dailyPresent{
		1:  {"first", "and a Partridge in a Pear Tree."},
		2:  {"second", "two Turtle Doves, "},
		3:  {"third", "three French Hens, "},
		4:  {"fourth", "four Calling Birds, "},
		5:  {"fifth", "five Gold Rings, "},
		6:  {"sixth", "six Geese-a-Laying, "},
		7:  {"seventh", "seven Swans-a-Swimming, "},
		8:  {"eighth", "eight Maids-a-Milking, "},
		9:  {"ninth", "nine Ladies Dancing, "},
		10: {"tenth", "ten Lords-a-Leaping, "},
		11: {"eleventh", "eleven Pipers Piping, "},
		12: {"twelfth", "twelve Drummers Drumming, "},
	}
	verse := "On the %v day of Christmas my true love gave to me: %v"
	if n == 1 {
		return fmt.Sprintf(verse, dailyPresents[n].day, "a Partridge in a Pear Tree.")
	}

	presents := make([]string, n)
	for i := n; i >= 1; i-- {
		presents[n-i] = dailyPresents[i].present
	}
	return fmt.Sprintf(verse, dailyPresents[n].day, strings.Join(presents, ""))
}

// Song gives the `The Twelve Days of Christmas` song
func Song() string {
	s := make([]string, 12)
	for d := 1; d <= 12; d++ {
		s[d-1] = Verse(d)
	}
	return strings.Join(s, "\n")
}
