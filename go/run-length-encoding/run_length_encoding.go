package encode

import (
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(i string) string {
	result := ""
	gr := group([]rune(i))
	for _, g := range gr {
		if len(g) == 1 {
			result += string(g[0])
			continue
		}
		result += strconv.Itoa(len(g)) + string(g[0])
	}
	return result
}

func RunLengthDecode(inp string) string {
	result := ""
	r := regexp.MustCompile(`\d*\D`)
	for _, i := range r.FindAllIndex([]byte(inp), -1) {
		c := inp[i[1]-1]
		n := 1
		if i[1]-i[0] != 1 {
			ns := inp[i[0] : i[1]-1]
			n, _ = strconv.Atoi(ns)
		}
		result += strings.Repeat(string(c), n)
	}
	return result
}

func group(inp []rune) [][]rune {
	grouped := [][]rune{}
	prev := []rune{}
	for i, r := range inp {
		if i == 0 {
			prev = []rune{r}
			continue
		}
		if r == prev[0] {
			prev = append(prev, r)
			if i == len(inp)-1 {
				grouped = append(grouped, prev)
			}
			continue
		}
		grouped = append(grouped, prev)
		prev = []rune{r}
		if i == len(inp)-1 {
			grouped = append(grouped, prev)
		}
	}
	return grouped
}
