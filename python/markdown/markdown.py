import re


class Markdown:
    def __init__(self, string):
        self.string = string
        self.segments = self.string.split("\n")

    def render(self):
        return f"{''.join(x.render() for x in self.parse())}"

    def parse(self):
        if not self.segments:
            return []
        parsed = []
        open_list = False
        for segment in self.segments:
            parsed_segment = self._parse_segment(segment)
            if not open_list and isinstance(parsed_segment, ListItem):
                open_list = True
                parsed_segment.first = True
                parsed.append(parsed_segment)
                continue
            if open_list and not isinstance(parsed_segment, ListItem):
                open_list = False
                parsed[-1].last = True
                parsed.append(parsed_segment)
                continue
            if open_list:
                parsed_segment.last = True
            parsed.append(parsed_segment)
        if open_list:
            parsed[-1].last = True
        return parsed

    def _parse_segment(self, segment):
        if header := self._header(segment):
            return header
        if list_item := self._list_item(segment):
            return list_item
        if bold := self._bold(segment):
            return bold
        if italic := self._italic(segment):
            return italic
        if text := self._text(segment):
            return text

    def _text(self, segment):
        if re.match(".+", segment):
            return Text(segment)

    def _header(self, segment):
        if not (match := re.match(r"^(#{1,6}) (.+)$", segment)):
            return
        return Header(len(match.group(1)), Text(match.group(2)))

    def _bold(self, segment):
        if not (match := re.match(r"^(.*)__([^_]+)__(.*)$", segment)):
            return
        result = [
            self._parse_segment(match.group(1)),
            Bold(self._parse_segment(match.group(2))),
            self._parse_segment(match.group(3)),
        ]
        return Text([x for x in result if x])

    def _italic(self, segment):
        if not (match := re.match(r"^(.*)_([^_]+)_(.*)$", segment)):
            return
        result = [
            self._parse_segment(match.group(1)),
            Italic(self._parse_segment(match.group(2))),
            self._parse_segment(match.group(3)),
        ]
        return Text([x for x in result if x])

    def _list_item(self, segment):
        if not (match := re.match(r"^( *)\* (.+)$", segment)):
            return
        return ListItem(self._parse_segment(match.group(2)))


class Header:
    def __init__(self, level, content):
        self.content = content
        self.level = level

    def __repr__(self):
        return f"Header({self.level}, {self.content})"

    def render(self):
        rendered = self.content.render(inside=True)
        return f"<h{self.level}>{rendered}</h{self.level}>"


class ListItem:
    def __init__(self, content):
        self.content = content
        self.first = False
        self.last = False

    def __repr__(self):
        return f"ListItem({self.content})"

    def _wrap_in_ul(self, rendered):
        if self.first:
            rendered = f"<ul>{rendered}"
        if self.last:
            rendered = f"{rendered}</ul>"
        return rendered

    def render(self, inside=False):
        if isinstance(self.content, list):
            rendered_list = (x.render(inside=True) for x in self.content)
            return self._wrap_in_ul(f"<li>{''.join(rendered_list)}</li>")
        if isinstance(self.content, str):
            return self._wrap_in_ul(f"<li>{self.content}</li>")
        return self._wrap_in_ul(f"<li>{self.content.render(inside=True)}</li>")


class Text:
    def __init__(self, content):
        self.content = content

    def __repr__(self):
        return f"Text({self.content})"

    def render(self, inside=False):
        if isinstance(self.content, list):
            rendered = "".join(x.render(inside=True) for x in self.content)
        elif isinstance(self.content, str):
            rendered = f"{self.content}"
        else:
            rendered = f"{self.content.render(inside=True)}"
        if inside:
            return rendered
        return f"<p>{rendered}</p>"


class Bold:
    def __init__(self, content):
        self.content = content

    def __repr__(self):
        return f"Bold({self.content})"

    def render(self, inside=False):
        if isinstance(self.content, list):
            rendered_list = (x.render(inside=True) for x in self.content)
            return f"<strong>{''.join(rendered_list)}</strong>"
        if isinstance(self.content, str):
            return f"<strong>{self.content}</strong>"
        return f"<strong>{self.content.render(inside=True)}</strong>"


class Italic:
    def __init__(self, content):
        self.content = content

    def __repr__(self):
        return f"Italic({self.content})"

    def render(self, inside=False):
        if isinstance(self.content, list):
            return f"<em>{''.join(x.render(inside=True) for x in self.content)}</em>"
        if isinstance(self.content, str):
            return f"<em>{self.content}</em>"
        return f"<em>{self.content.render(inside=True)}</em>"


def parse(content):
    markdown = Markdown(content)
    return markdown.render()


if __name__ == "__main__":
    tests = [
        "_This will be italic_",
        "This will be a paragraph",
        "__This will be bold__",
        "This will _be_ __mixed__",
        "# This will be an h1",
        "## This will be an h2",
        "###### This will be an h6",
        "* Item 1\n* Item 2",
        "# Header!\n* __Bold Item__\n* _Italic Item_",
        "# This is a header with # and * in the text",
        "* Item 1 with a # in the text\n* Item 2 with * in the text",
        "This is a paragraph with # and * in the text",
        "# Before L1\n* L1 1\n* L1 2\nAfter L1\n* L2 1\n* L2 2\n* L3 2\nAfter L2",
    ]
    for test in tests:
        m_test = Markdown(test)
        print(f"{m_test.segments=}\n{m_test.parse()=}\n{m_test.render()=}\n")
