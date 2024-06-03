package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/conikeec/markdown-parser/pkg/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: parser <markdown_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	document, err := parser.ParseMarkdownFile(filePath)
	if err != nil {
		fmt.Printf("Error parsing Markdown file: %v\n", err)
		os.Exit(1)
	}

	pprint, _ := json.MarshalIndent(document, "", "\t")
	fmt.Println(string(pprint))
}
