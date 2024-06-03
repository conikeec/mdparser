# Building and Running the Program
The Markdown Parser is a command-line tool that parses Markdown files and extracts the document structure, including sections, content, and code blocks. To build and run the program, follow these steps:

Open a terminal and navigate to the markdown-parser directory.

## Run

```go build ./cmd/parser```

```./parser <markdown_file>```

For example, if you have a Markdown file named example.md in the examples directory, you can run:
Copy code ```./parser examples/[filename].md```
The program will output the parsed Markdown document structure, displaying the sections, content, and code blocks.
Feel free to explore the provided examples in the examples directory to see how the Markdown Parser handles different Markdown syntax and structures.

## Purpose 
The parser serves to maintain a structural interchange between LLM invocations akin to message passing and processing. 
LLMs are non-determinstic with their output format despite of being instructed.
If this parser is being leveraged to strcuture the out of a LLM, ensure that the prompt is augmented with instructions to format the output

Example
```md
            **Output Format**: 
            Please provide the response in Markdown format with the following structure:
            # Fix
            ```{programming_language}
            ...
            ```
            # Dependency Update
            ```{programming_language}
            ...
            ```
            # Import
            ```{programming_language}
            ...
            ```
            # Notes
            ...

            **Format Instructions**: As indicated in the Output Format, format the response output as a markdown object. The object should have five fields: Fix (output created only in code-block section), Dependency Update (output created only in code-block section and no comments, headers in this section), Import (output created only in code-block section) and Notes (created only in content section). No other comments, headers and descriptions are allowed in the output.

            **Instructions**
            ....
            12. Add all notes in the *Notes* section of the Output Format above. Review the notes for repetition. Remove any duplicate or repetitive notes. Aim for a concise set of unique, insightful notes.


```