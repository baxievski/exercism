package encode

import (
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(i string) string {
	result := []rune{}
	gr := group([]rune(i))
	for _, g := range gr {
		if len(g) == 1 {
			result = append(result, g[0])
			continue
		}
		c := []rune(strconv.Itoa(len(g)))
		c = append(c, g[0])
		result = append(result, c...)
	}
	return string(result)
}

func RunLengthDecode(inp string) string {
	result := []rune{}
	r := regexp.MustCompile(`\d*\D`)
	for _, i := range r.FindAllIndex([]byte(inp), -1) {
		c := inp[i[1]-1]
		n := 1
		if i[1]-i[0] != 1 {
			ns := inp[i[0] : i[1]-1]
			n, _ = strconv.Atoi(ns)
		}
		result = append(result, []rune(strings.Repeat(string(c), n))...)
	}
	return string(result)
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
