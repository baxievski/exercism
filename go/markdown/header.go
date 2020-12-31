package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

type header struct {
	content []renderer
	level   int
}

func (h header) render() string {
	sb := strings.Builder{}
	for _, c := range h.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<h%v>%s</h%v>", h.level, sb.String(), h.level)
}

func parseHeader(ln string) renderer {
	sm := reSubMatch(regexp.MustCompile(`^(?P<H>#{1,6}) (?P<HContent>.*)$`), ln)
	if len(sm) == 0 {
		return nil
	}
	return header{
		level:   len(sm["H"]),
		content: []renderer{text{content: sm["HContent"]}},
	}
}
