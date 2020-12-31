package markdown

// Source: exercism/problem-specifications
// Commit: 9713d52 markdown: Apply new "input" policy
// Problem Specifications Version: 1.2.0

// Markdown is a shorthand for creating HTML from text strings.

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "parses normal text as a paragraph",
		input:       "This will be a paragraph",
		expected:    "<p>This will be a paragraph</p>",
	},
	{
		description: "parsing italics",
		input:       "_This will be italic_",
		expected:    "<p><em>This will be italic</em></p>",
	},
	{
		description: "parsing bold text",
		input:       "__This will be bold__",
		expected:    "<p><strong>This will be bold</strong></p>",
	},
	{
		description: "mixed normal, italics and bold text",
		input:       "This will _be_ __mixed__",
		expected:    "<p>This will <em>be</em> <strong>mixed</strong></p>",
	},
	{
		description: "with h1 header level",
		input:       "# This will be an h1",
		expected:    "<h1>This will be an h1</h1>",
	},
	{
		description: "with h2 header level",
		input:       "## This will be an h2",
		expected:    "<h2>This will be an h2</h2>",
	},
	{
		description: "with h6 header level",
		input:       "###### This will be an h6",
		expected:    "<h6>This will be an h6</h6>",
	},
	{
		description: "with symbols in header",
		input:       "This is a paragraph with # and * in the text",
		expected:    "<p>This is a paragraph with # and * in the text</p>",
	},
	{
		description: "unordered lists",
		input:       "* Item 1\n* Item 2",
		expected:    "<ul><li>Item 1</li><li>Item 2</li></ul>",
	},
	{
		description: "multiple unordered lists",
		input:       "# Header\n* Item 1\n* Item 2\n## Header 2\n* Item A\n * Item B",
		expected:    "<h1>Header</h1><ul><li>Item 1</li><li>Item 2</li></ul><h2>Header 2</h2><ul><li>Item A</li><li>Item B</li></ul>",
	},
	{
		description: "With a little bit of everything",
		input:       "# Header!\n* __Bold Item__\n* _Italic Item_",
		expected:    "<h1>Header!</h1><ul><li><strong>Bold Item</strong></li><li><em>Italic Item</em></li></ul>",
	},
}
