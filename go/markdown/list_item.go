package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

type listItem struct {
	content []renderer
}

func (li listItem) render() string {
	sb := strings.Builder{}
	for _, c := range li.content {
		sb.WriteString(c.render())
	}

	return fmt.Sprintf("<li>%s</li>", sb.String())
}

func parseListItem(line string) renderer {
	l := regexp.MustCompile(`^(?P<LiIndentation> *)\* (?P<Li>.*)$`)
	submatch := reSubMatch(l, line)
	if len(submatch) == 0 {
		return nil
	}
	return listItem{content: parseSegment(submatch["Li"])}
}
