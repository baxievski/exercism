package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type searchResult struct {
	text     string
	lineNum  int
	fileName string
}

type searchResults []searchResult
type options struct {
	ignoreCase    bool
	invert        bool
	wholeLine     bool
	lineNumbers   bool
	fileNames     bool
	multipleFiles bool
}

type grep struct {
	opts    options
	results []searchResult
}

func Search(pattern string, flags []string, filePaths []string) []string {
	// g := grep{}
	results := searchResults{}
	opts := parseOptions(flags)
	for i, path := range filePaths {
		if i >= 1 && !opts.multipleFiles {
			opts.multipleFiles = true
		}
		m, err := searchFile(pattern, path, opts)
		if err != nil {
			continue
		}
		results = append(results, m...)
	}
	return results.format(opts)
}

func searchFile(pattern string, path string, opts options) ([]searchResult, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	if opts.wholeLine {
		pattern = fmt.Sprintf("^%v$", pattern)
	}
	if opts.ignoreCase {
		pattern = fmt.Sprintf("(?i)%v", strings.ToLower(pattern))
	}
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	results := searchResults{}
	ln := 0
	for {
		ln++
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line = strings.TrimSuffix(line, "\n")
		if reg.MatchString(line) != opts.invert {
			results = append(results, searchResult{text: line, lineNum: ln, fileName: path})
			if opts.fileNames {
				break
			}
		}
	}
	return results, nil
}

func (g *grep) parseOptions(flags []string) {
	for _, v := range flags {
		switch v {
		case "-l":
			// g.opts.fileNames
			g.opts.fileNames = true
		case "-n":
			g.opts.lineNumbers = true
		case "-i":
			g.opts.ignoreCase = true
		case "-v":
			g.opts.invert = true
		case "-x":
			g.opts.wholeLine = true
		}
	}
}
func parseOptions(flags []string) options {
	o := options{}
	for _, v := range flags {
		switch v {
		case "-l":
			o.fileNames = true
		case "-n":
			o.lineNumbers = true
		case "-i":
			o.ignoreCase = true
		case "-v":
			o.invert = true
		case "-x":
			o.wholeLine = true
		}
	}
	return o
}

func (sr *searchResults) format(opts options) []string {
	r := []string{}
	c := ""
	for _, m := range *sr {
		if opts.fileNames {
			r = append(r, m.fileName)
			continue
		}
		c = m.text
		if opts.lineNumbers {
			c = fmt.Sprintf("%v:%s", m.lineNum, c)
		}
		if opts.multipleFiles {
			c = fmt.Sprintf("%s:%s", m.fileName, c)
		}
		r = append(r, c)
	}
	return r
}
