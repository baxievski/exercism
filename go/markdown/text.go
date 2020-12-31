package markdown

type text struct {
	content string
}

func (t text) render() string {
	return t.content
}
