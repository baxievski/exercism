package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

type paragraph struct {
	content []renderer
}

func (p paragraph) render() string {
	sb := strings.Builder{}
	for _, c := range p.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<p>%s</p>", sb.String())
}

func parseParagraph(line string) renderer {
	if regexp.MustCompile(`^#{1,6}.*`).Match([]byte(line)) {
		return nil
	}
	l := regexp.MustCompile(`^\s*\* .*`)
	if l.Match([]byte(line)) {
		return nil
	}
	return paragraph{content: parseSegment(line)}
}
