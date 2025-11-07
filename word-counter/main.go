package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	countLines := false

	if len(os.Args) > 1 {
		if os.Args[1] == "-l" {
			countLines = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	total := 0

	if countLines {
		fmt.Println("Write your text (Counting Lines):")
	} else {
		fmt.Println("Write your text (Counting Words):")
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.ToLower(strings.TrimSpace(line)) == "exit" {
			break
		}

		if countLines {
			total++

		} else {
			words := strings.Fields(line)
			total += len(words)
		}
	}

	if countLines {

		fmt.Printf("Lines: %d\n ", total)

	} else {
		fmt.Printf("Words: %d\n ", total)

	}
}
