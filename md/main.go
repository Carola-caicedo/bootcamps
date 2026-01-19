package main

import (
	"flag"
	"fmt"
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

	if err := run(*inflag, *outflag); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func run(in, out string) error {
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

	var filename string
	if out != "" {
		filename = out + ".html"
	} else {
		base := filepath.Base(in)
		ext := filepath.Ext(base)
		name := base[:len(base)-len(ext)]
		filename = name + ".html"
	}

	content := []byte(header + string(body) + footer)

	return saveHTML(filename, content)
}

func parseContent(input []byte) ([]byte, error) {
    output := blackfriday.Run(input)
    safeHTML := bluemonday.UGCPolicy().SanitizeBytes(output)
    return safeHTML, nil
}


func saveHTML(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
