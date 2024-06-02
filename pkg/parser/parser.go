package parser

import (
	"fmt"
	"os"
	"strings"
)

type MarkdownDocument struct {
	Sections []Section
}

type Section struct {
	Title      string
	Content    string
	CodeBlocks []CodeBlock
}

type CodeBlock struct {
	Language string
	Code     string
}

func ParseMarkdownString(markdown string) (MarkdownDocument, error) {
	if strings.TrimSpace(markdown) == "" {
		return MarkdownDocument{}, fmt.Errorf("empty markdown document")
	}

	var sections []Section
	lines := strings.Split(markdown, "\n")

	var currentSection Section
	var inCodeBlock bool
	var currentCodeBlock CodeBlock

	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				currentSection.CodeBlocks = append(currentSection.CodeBlocks, currentCodeBlock)
				currentCodeBlock = CodeBlock{}
			} else {
				language := strings.TrimSpace(strings.TrimPrefix(line, "```"))
				currentCodeBlock.Language = language
			}
			inCodeBlock = !inCodeBlock
		} else if inCodeBlock {
			currentCodeBlock.Code += line + "\n"
		} else if strings.HasPrefix(line, "#") {
			if currentSection.Title != "" {
				sections = append(sections, currentSection)
				currentSection = Section{}
			}
			currentSection.Title = strings.TrimSpace(strings.TrimPrefix(line, "#"))
		} else {
			currentSection.Content += line + "\n"
		}
	}

	if currentSection.Title != "" {
		sections = append(sections, currentSection)
	}

	return MarkdownDocument{Sections: sections}, nil
}

func ParseMarkdownFile(filePath string) (MarkdownDocument, error) {
	// Read the Markdown file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return MarkdownDocument{}, fmt.Errorf("failed to read file: %v", err)
	}

	// Parse the Markdown content
	return ParseMarkdownString(string(content))
}
