package markdown

import (
	"fmt"
	"strings"
)

type renderer interface {
	render() string
}

type header struct {
	content []renderer
	level   int
}

type listItem struct {
	content []renderer
}

type unorderedList struct {
	content []renderer
}

type text struct {
	content []renderer
	nested  bool
}

type bold struct {
	content []renderer
}

type italic struct {
	content []renderer
}

func (t *text) render() string {
	sb := strings.Builder{}
	for _, c := range t.content {
		sb.WriteString(c.render())
	}
	if t.nested {
		return sb.String()
	}
	return fmt.Sprintf("<p>%s</p>", sb.String())

}

func (b *bold) render() string {
	sb := strings.Builder{}
	for _, c := range b.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<strong>%s</strong>", sb.String())
}

func (i *italic) render() string {
	sb := strings.Builder{}
	for _, c := range i.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<em>%s</em>", sb.String())
}

func (h *header) render() string {
	sb := strings.Builder{}
	for _, c := range h.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<h%v>%s</h%v>", h.level, sb.String(), h.level)
}

func (li *listItem) render() string {
	sb := strings.Builder{}
	for _, c := range li.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<li>%s</li>", sb.String())
}

func (ul *unorderedList) render() string {
	sb := strings.Builder{}
	for _, c := range ul.content {
		sb.WriteString(c.render())
	}
	return fmt.Sprintf("<ul>%s</ul>", sb.String())
}

// TODO: use AST

func getNode(s string) renderer {
	// lines := strings.Split(s, "\n")
	// nodes := []renderer{}
	// for i, line := range lines {

	// }

	return &text{}
}

// Render translates markdown to HTML
func Render(markdown string) string {
	header := 0
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	pos := 0
	list := 0
	html := ""
	for {
		char := markdown[pos]
		if char == '#' {
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
			pos++
			continue
		}
		if char == '*' {
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos += 2
			continue
		}
		if char == '\n' {
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
			pos++
			continue
		}
		html += string(char)
		pos++
		if pos >= len(markdown) {
			break
		}
	}
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"

}
