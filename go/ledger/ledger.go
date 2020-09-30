package ledger

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Locale holds the localization templates
type Locale struct {
	name                     string
	dateFormat               string
	header                   string
	positiveCurrencyFormat   string
	negativeCurrencyFormat   string
	decimalSymbol            string
	thousandsSeparatorSymbol string
}

var locales = map[string]Locale{
	"en-US": {
		"en-US",
		"01/02/2006",
		"Date       | Description               | Change",
		"%s%v ",
		"(%s%v)",
		".",
		",",
	},
	"nl-NL": {
		"nl-NL",
		"02-01-2006",
		"Datum      | Omschrijving              | Verandering",
		"%s %v ",
		"%s %v-",
		",",
		".",
	},
}

var currencies = map[string]string{
	"USD": "$",
	"EUR": "€",
}

// Entry represents an entry in a transaction log
// Date is in YYYY-MM-DD
type Entry struct {
	Date        string
	Description string
	Change      int
}

func (e *Entry) date(l Locale) (string, error) {
	t, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return "", fmt.Errorf("invalid date '%v'", e.Date)
	}
	return t.Format(l.dateFormat), nil
}

func (e *Entry) description(l int) string {
	d := []rune(e.Description)
	if len(d) < l-3 {
		return e.Description
	}
	return string(d[:l]) + "..."
}

func (e *Entry) change(l Locale, currencySymbol string) string {
	c := float64(e.Change) / 100
	f := l.positiveCurrencyFormat
	if e.Change < 0 {
		c = 0 - c
		f = l.negativeCurrencyFormat
	}
	buf := &bytes.Buffer{}
	parts := strings.Split(strconv.FormatFloat(c, 'f', 2, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.WriteString(l.thousandsSeparatorSymbol)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.WriteString(l.thousandsSeparatorSymbol)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.WriteString(l.decimalSymbol)
		buf.WriteString(parts[1])
	}

	return fmt.Sprintf(f, currencySymbol, buf.String())
}

// Ledger is a list of transactions
type Ledger []Entry

func (l Ledger) sort() Ledger {
	entries := make([]Entry, 0, len(l))
	for _, e := range l {
		entries = append(entries, e)
	}
	sort.Slice(entries, func(i, j int) bool {
		tI, _ := time.Parse("2006-01-01", entries[i].Date)
		tJ, _ := time.Parse("2006-01-01", entries[j].Date)
		if !tI.Equal(tJ) {
			return tI.Before(tJ)
		}
		if entries[i].Description != entries[j].Description {
			return entries[i].Description < entries[j].Description
		}
		return entries[i].Change < entries[j].Change
	})
	return entries
}

// FormatLedger gives a localized representation of the transaction log
func FormatLedger(currency string, locale string, entries Ledger) (string, error) {
	currencySymbol, curOk := currencies[currency]
	if !curOk {
		return "", fmt.Errorf("invalid currency '%v'", currency)
	}
	loc, locOk := locales[locale]
	if !locOk {
		return "", fmt.Errorf("invalid locale '%v'", locale)
	}
	sb := strings.Builder{}
	sb.WriteString(loc.header)
	sb.WriteString("\n")
	for _, e := range entries.sort() {
		date, err := e.date(loc)
		if err != nil {
			return "", err
		}
		sb.WriteString(date)
		sb.WriteString(" | ")
		sb.WriteString(fmt.Sprintf("%-25v", e.description(22)))
		sb.WriteString(" | ")
		sb.WriteString(fmt.Sprintf("%13v", e.change(loc, currencySymbol)))
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
