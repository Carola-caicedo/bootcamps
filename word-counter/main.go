package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//function that performs the counting process
func countString(input string, countLines bool) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(line), "exit") {
			break
		}

		if countLines {
			total++
		} else {
			words := strings.Fields(line)
			total += len(words)
		}
	}
	return total
}


//main function
func main() {

	countLines := false

	if len(os.Args) > 1 {
		if os.Args[1] == "-l" {
			countLines = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder

	if countLines {
		fmt.Println("Write your text (Counting Lines):")
	} else {
		fmt.Println("Write your text (Counting Words):")
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(line), "exit") {
			break
		}

		input.WriteString(line + "\n")

	}

	// Call the counting function
	total := countString(input.String(), countLines)

	if countLines {

		fmt.Printf("Lines: %d\n ", total)

	} else {
		fmt.Printf("Words: %d\n ", total)

	}
}
