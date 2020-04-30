import re


class Markdown:
    def __init__(self, string):
        super().__init__()
        self.string = string
        self.segments = self.string.split("\n")
    
    def render(self):
        return f"{''.join(x.render() for x in self.parse())}"

    def parse(self):
        parsed = []
        for i, segment in enumerate(self.segments):
            if i == 0:
                parsed.append(self._parse_segment(segment))
                continue
            if isinstance(parsed[i-1], UnorderedList):
                _ = self._parse_segment(segment, unordered_list=parsed[i-1])
                parsed.append(Empty())
                continue
            parsed.append(self._parse_segment(segment))
        return parsed
    
    def _parse_segment(self, segment, unordered_list=None):
        if (header := self._header(segment)) is not None:
            return header
        if (list_item := self._list_item(segment)) is not None:
            if unordered_list is None:
                unordered_list = UnorderedList()
            unordered_list.content.append(list_item)
            return unordered_list
        if (bold := self._bold(segment)) is not None:
            return bold
        if (italic := self._italic(segment)) is not None:
            return italic
        if (text := self._text(segment)) is not None:
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
        result = Bold(self._parse_segment(match.group(2)))
        if match.group(1):
            result = [self._parse_segment(match.group(1)), result]
        if match.group(3):
            if not isinstance(result, list):
                result = [result, ]
            result.append(self._parse_segment(match.group(3)))
        return Text(result)

    def _italic(self, segment):
        if not (match := re.match(r"^(.*)_([^_]+)_(.*)$", segment)):
            return
        result = Italic(self._parse_segment(match.group(2)))
        if match.group(1):
            result = [self._parse_segment(match.group(1)), result]
        if match.group(3):
            if not isinstance(result, list):
                result = [result, ]
            result.append(self._parse_segment(match.group(3)))
        return Text(result)

    def _list_item(self, segment):
        if not (match := re.match(r"^( *)\* (.+)$", segment)):
            return
        return ListItem(self._parse_segment(match.group(2)))


class Header:
    def __init__(self, level, content):
        super().__init__()
        self.content = content
        self.level = level

    def __repr__(self):
        return f"Header({self.level}, {self.content})"
    
    def render(self):
        rendered = self.content.render(inside=True)
        return f"<h{self.level}>{rendered}</h{self.level}>"


class ListItem:
    def __init__(self, content):
        super().__init__()
        self.content = content

    def __repr__(self):
        return f"ListItem({self.content})"

    def render(self, inside=False):
        if isinstance(self.content, list):
            return f"<li>{''.join(x.render(inside=True) for x in self.content)}</li>"
        if isinstance(self.content, str):
            return f"<li>{self.content}</li>"
        return f"<li>{self.content.render(inside=True)}</li>"


class UnorderedList:
    def __init__(self):
        super().__init__()
        self.content = []

    def __repr__(self):
        return f"UnorderedList({self.content})"

    def render(self):
        print(f"UnorderedList.render(): {self}")
        if self.content:
            return f"<ul>{''.join(x.render(inside=True) for x in self.content)}</ul>"


class Text:
    def __init__(self, content):
        super().__init__()
        self.content = content

    def __repr__(self):
        return f"Text({self.content})"

    def render(self, inside=False):
        if isinstance(self.content, list):
            rendered =  "".join(x.render(inside=True) for x in self.content)
        elif isinstance(self.content, str):
            rendered = f"{self.content}"
        else:
            rendered = f"{self.content.render(inside=True)}"
        if inside:
            return rendered
        return f"<p>{rendered}</p>"


class Empty:
    def __init__(self):
        super().__init__()

    def __repr__(self):
        return f"Empty()"

    def render(self, inside=False):
        return f""


class Bold:
    def __init__(self, content):
        super().__init__()
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
        super().__init__()
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
        # "This will be a paragraph",
        # "__This will be bold__",
        # "This will _be_ __mixed__",
        # "# This will be an h1",
        # "## This will be an h2",
        "###### This will be an h6",
        # "* Item 1\n* Item 2",
        # "# Header!\n* __Bold Item__\n* _Italic Item_",
        # "# This is a header with # and * in the text",
        # "* Item 1 with a # in the text\n* Item 2 with * in the text",
        # "This is a paragraph with # and * in the text",
        "# Start a list\n* Item 1\n* Item 2\nEnd a list"
    ]
    for test in tests:
        m_test = Markdown(test)
        print(f"{m_test.segments=}\n{m_test.parse()=}\n{m_test.render()=}\n")
