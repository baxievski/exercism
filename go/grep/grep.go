package grep

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type match struct {
	text     string
	lineNum  int
	fileName string
}

type matchFn func(a, b string) bool

func Search(pattern string, flags []string, files []string) []string {
	matched := []match{}
	opts := parseOptions(flags)
	fn := func(a, b string) bool { return strings.Contains(a, b) }
	if opts["-i"] {
		fn = func(a, b string) bool { return strings.Contains(strings.ToLower(a), strings.ToLower(b)) }
	}
	if opts["-x"] {
		fn = func(a, b string) bool { return a == b }
	}
	for _, f := range files {
		m, err := searchFile(pattern, f, fn)
		if err != nil {
			continue
		}
		matched = append(matched, m...)
	}
	r := []string{}
	for _, m := range matched {
		r = append(r, m.text)
	}
	return r
}

func searchFile(pattern string, path string, fn matchFn) ([]match, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bf := bufio.NewReader(f)

	m := []match{}
	ln := 0
	for {
		l, err := bf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if fn(l, pattern) {
			m = append(m, match{text: strings.TrimSuffix(l, "\n"), lineNum: ln, fileName: path})
		}
		ln++
	}
	return m, nil
}

func parseOptions(flags []string) map[string]bool {
	o := make(map[string]bool)
	for _, v := range flags {
		switch v {
		case "-n":
			o["n"] = true
		case "-l":
			o["l"] = true
		case "-i":
			o["i"] = true
		case "-v":
			o["v"] = true
		case "-x":
			o["x"] = true
		}
	}
	return o
}

// func formatResullts(matches []match, flags []string) []string {
// 	var fN, fL, fI, fV, fX bool
// 	for _, fl := range flags {
// 		if fl == "-n" {
// 			fN = true
// 			continue
// 		}
// 		if fl == "-l" {
// 			fL = true
// 			continue
// 		}
// 		if fl == "-l" {
// 			fL = true
// 			continue
// 		}
// 		if fl == "-i" {
// 			fI = true
// 			continue
// 		}
// 		if fl == "-x" {
// 			fX = true
// 			continue
// 		}
// 	}
// 	templ := "%v"
// 	if fN {
// 		templ = "%v:%v"
// 	}
// }
