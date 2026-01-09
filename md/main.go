package main

import (
	"flag"
	"fmt"
	"os"
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
	outflag := flag.String("out", "", "HTML filename")
	flag.Parse()

	if err := run(*outflag); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func run(out string) error {
	// Check obligatory flag
	if out == "" {
		return fmt.Errorf("the flag (-out) is obligatory ")
	}

	filename := out + ".html"

	content := []byte(header + footer)

	return saveHTML(filename, content)
}

func saveHTML(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
