# Building and Running the Program
The Markdown Parser is a command-line tool that parses Markdown files and extracts the document structure, including sections, content, and code blocks. To build and run the program, follow these steps:

Open a terminal and navigate to the markdown-parser directory.

Run ```go build ./cmd/parser to build the binary. This will generate an executable named parser in the current directory.```
Run ```./parser <markdown_file> to parse a Markdown file. Replace <markdown_file> with the path to your Markdown file.```

For example, if you have a Markdown file named example.md in the examples directory, you can run:
Copy code ```./parser examples/[filename].md```
The program will output the parsed Markdown document structure, displaying the sections, content, and code blocks.
Feel free to explore the provided examples in the examples directory to see how the Markdown Parser handles different Markdown syntax and structures.