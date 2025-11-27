package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// function that performs the counting process
func countString(input string, countLines bool, countBytes bool) int {

	if countLines {

		// Count the number of lines
		scanner := bufio.NewScanner(strings.NewReader(input))
		total := 0

		for scanner.Scan() {
			line := scanner.Text()

			if strings.EqualFold(strings.TrimSpace(line), "exit") {
				break
			}

			total++
		}

		return total
	} else if countBytes {
		// Count the number of bytes
		return len([]byte(input))

	} else {

		// Count the number of words
		scanner := bufio.NewScanner(strings.NewReader(input))
		total := 0
		for scanner.Scan() {
			line := scanner.Text()

			if strings.EqualFold(strings.TrimSpace(line), "exit") {

				break
			}

			words := strings.Fields(line)
			total += len(words)
		}
		return total
	}

}

// main function
func main() {

	// Define flags
	//countLines
	countLines := flag.Bool("l", false, "Count lines")
	//countBytes
	countBytes := flag.Bool("b", false, "Count bytes")
	//countWords (default)

	//Show status before parsing
	fmt.Printf("Before parsing - Lines: %t, Bytes: %t\n", *countLines, *countBytes)

	// Parse the command line flags
	flag.Parse()

	// Show status after parsing
	fmt.Printf("After parsing - Lines: %t, Bytes: %t\n", *countLines, *countBytes)

	// Define the mode of counting
	var mode string

	if *countLines {
		mode = "lines"
	} else if *countBytes {
		mode = "bytes"
	} else {
		mode = "words"
	}

	// Print the mode of counting
	fmt.Printf("The program will count %s.\n", mode)

	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder

	fmt.Printf("Write your text (Counting %s):", mode)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(line), "exit") {
			break
		}

		input.WriteString(line + "\n")

	}

	// Call the counting function
	total := countString(input.String(), *countLines, *countBytes)

	switch mode {
	case "lines":
		fmt.Printf("Lines: %d\n ", total)
	case "bytes":
		fmt.Printf("bytes: %d\n ", total)
	default:
		fmt.Printf("Words: %d\n ", total)
	}

}
