package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const header = `<!DOCTYPE html>
<html>
    <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8" />
    <title>Markdown Preview Tool</title>
    </head>
    <body>
`

const footer = `
    </body>
</html>
`

func main() {
	// define flag
	inflag := flag.String("in", "", "Markdown filename")
	outflag := flag.String("out", "", "HTML filename")
	flag.Parse()

	if err := run(*inflag, *outflag, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func run(in, out string, writer io.Writer) error {
	// Check obligatory flag
	if in == "" {
		return fmt.Errorf("the flag (-in) is obligatory")
	}

	input, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("cannot read the file: %v", err)
	}

	body, err := parseContent(input)
	if err != nil {
		return err
	}

	if err := os.MkdirAll("md", 0755); err != nil {
		return fmt.Errorf("cannot created directory: %v", err)
	}

	var filename string
	if out != "" {
		filename = filepath.Join("md", out+".html")
	} else {
		tempFile, err := os.CreateTemp("md", "md*.html")
		if err != nil {
			return fmt.Errorf("cannot created file temporary: %v", err)
		}

		tempFile.Close()
		filename = tempFile.Name()

	}

	content := []byte(header + string(body) + footer)

	if err := saveHTML(filename, content); err != nil {
		return err
	}

	fmt.Fprintln(writer, filename)
	return nil
}

func parseContent(input []byte) ([]byte, error) {
	output := blackfriday.Run(input, blackfriday.WithExtensions(blackfriday.CommonExtensions))
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(output)
	return safeHTML, nil
}

func saveHTML(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
