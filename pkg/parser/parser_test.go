package parser

import (
	"reflect"
	"testing"
)

func TestParseMarkdownString(t *testing.T) {
	// Test case 1: Valid Markdown string
	markdown := "# Section 1\nContent 1\n```python\nprint('Hello, World!')\n```\n\n# Section 2\nContent 2"
	expected := MarkdownDocument{
		Sections: []Section{
			{
				Title:   "Section 1",
				Content: "Content 1\n",
				CodeBlocks: []CodeBlock{
					{
						Language: "python",
						Code:     "print('Hello, World!')\n",
					},
				},
			},
			{
				Title:   "Section 2",
				Content: "Content 2",
			},
		},
	}
	result, err := ParseMarkdownString(markdown)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, result)
	}

	// Test case 2: Empty Markdown string
	markdown = ""
	_, err = ParseMarkdownString(markdown)
	if err == nil {
		t.Error("Expected an error for empty Markdown string, but got nil")
	}
}
