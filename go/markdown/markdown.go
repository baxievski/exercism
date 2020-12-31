package markdown

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type renderer interface {
	render() string
}

// Render translates markdown to HTML
func Render(input string) string {
	sb := strings.Builder{}
	insideList := false
	for _, n := range getNodes(strings.NewReader(input)) {
		if _, li := n.(listItem); li && !insideList {
			insideList = true
			sb.WriteString("<ul>")
		}
		if _, li := n.(listItem); !li && insideList {
			insideList = false
			sb.WriteString("</ul>")
		}
		sb.WriteString(n.render())
	}
	if insideList {
		sb.WriteString("</ul>")
	}
	return sb.String()
}

func reSubMatch(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatch := make(map[string]string)

	if len(match) != 0 {
		for i, name := range r.SubexpNames() {
			if i != 0 {
				subMatch[name] = match[i]
			}
		}
	}

	return subMatch
}

func parseSegment(segment string) []renderer {
	result := []renderer{}

	b := regexp.MustCompile(`^(?P<PreB>.*)__(?P<B>[^_]+)__(?P<PostB>.*)$`)
	if submatch := reSubMatch(b, segment); len(submatch) != 0 {
		if preBold := submatch["PreB"]; len(preBold) != 0 {
			result = append(result, parseSegment(preBold)...)
		}
		result = append(result, bold{content: []renderer{text{content: submatch["B"]}}})
		if postBold := submatch["PostB"]; len(postBold) != 0 {
			result = append(result, parseSegment(postBold)...)
		}
		return result
	}

	i := regexp.MustCompile(`^(?P<PreI>.*)_(?P<I>[^_]+)_(?P<PostI>.*)$`)
	if submatch := reSubMatch(i, segment); len(submatch) != 0 {
		if preItalic := submatch["PreI"]; len(preItalic) != 0 {
			result = append(result, parseSegment(preItalic)...)
		}
		result = append(result, italic{content: []renderer{text{content: submatch["I"]}}})
		if postItalic := submatch["PostI"]; len(postItalic) != 0 {
			result = append(result, parseSegment(postItalic)...)
		}
		return result
	}

	t := regexp.MustCompile(`(?P<T>.+)`)
	if submatch := reSubMatch(t, segment); len(submatch) != 0 {
		return []renderer{text{content: submatch["T"]}}
	}

	return []renderer{}
}

func parseLine(line string) renderer {
	if h := parseHeader(line); h != nil {
		return h
	}

	if l := parseListItem(line); l != nil {
		return l
	}

	if p := parseParagraph(line); p != nil {
		return p
	}

	return nil
}

func getNodes(input io.Reader) []renderer {
	parsed := []renderer{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		n := parseLine(scanner.Text())
		parsed = append(parsed, n)
	}
	return parsed
}
