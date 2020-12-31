package markdown

import (
	"fmt"
	"strings"
)

type italic struct {
	content []renderer
}

func (i italic) render() string {
	sb := strings.Builder{}
	for _, c := range i.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<em>%s</em>", sb.String())
}
