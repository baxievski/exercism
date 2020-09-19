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

// Search for lines in files matching a regular expression
func Search(pattern string, flags []string, filePaths []string) []string {
	g := grep{}
	g.parseOptions(flags)
	for i, path := range filePaths {
		if i >= 1 && !g.opts.multipleFiles {
			g.opts.multipleFiles = true
		}
		if g.searchFile(pattern, path) != nil {
			continue
		}
	}
	return g.print()
}

func (g *grep) searchFile(pattern string, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if g.opts.wholeLine {
		pattern = fmt.Sprintf("^%v$", pattern)
	}
	if g.opts.ignoreCase {
		pattern = fmt.Sprintf("(?i)%v", pattern)
	}

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)
	ln := 0
	for {
		ln++
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line = strings.TrimSuffix(line, "\n")
		if reg.MatchString(line) != g.opts.invert {
			g.addResult(line, ln, path)
			if g.opts.fileNames {
				break
			}
		}
	}
	return nil
}

func (g *grep) addResult(text string, lineNum int, fileName string) {
	g.results = append(g.results, searchResult{text: text, lineNum: lineNum, fileName: fileName})
}

func (g *grep) parseOptions(flags []string) {
	for _, v := range flags {
		switch v {
		case "-l":
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

func (g *grep) print() []string {
	out := []string{}
	tx := ""
	for _, m := range g.results {
		if g.opts.fileNames {
			out = append(out, m.fileName)
			continue
		}
		tx = m.text
		if g.opts.lineNumbers {
			tx = fmt.Sprintf("%v:%s", m.lineNum, tx)
		}
		if g.opts.multipleFiles {
			tx = fmt.Sprintf("%s:%s", m.fileName, tx)
		}
		out = append(out, tx)
	}
	return out
}
