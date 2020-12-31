package markdown

import (
	"fmt"
	"strings"
)

type bold struct {
	content []renderer
}

func (b bold) render() string {
	sb := strings.Builder{}
	for _, c := range b.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<strong>%s</strong>", sb.String())
}
