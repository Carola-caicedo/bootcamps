package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// func that reads user input
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line is the exit command
		if strings.EqualFold(strings.TrimSpace(line), "exit") {
			break
		}
		input.WriteString(line + "\n")

	}
	return input.String()

}

// function that does the counting
func countString(input string, countLines bool, countBytes bool) int {

	// Check if the line is the exit command
	if strings.EqualFold(strings.TrimSpace(input), "exit") {
		return 0
	}

	//Priority
	if countLines {
		// Count the number of Lines
		scanner := bufio.NewScanner(strings.NewReader(input))
		total := 0

		for scanner.Scan() {
			scanner.Text()
			total++
		}

		return total
	}

	if countBytes {
		// Count the number of Bytes
		return len([]byte(input))
	}

	// Count words by default
	scanner := bufio.NewScanner(strings.NewReader(input))
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		total += len(words)
	}
	return total

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
	// fmt.Printf("Before parsing - Lines: %t, Bytes: %t\n", *countLines, *countBytes)

	// processes the flags arguments in line
	flag.Parse()

	// Show status after parsing
	// fmt.Printf("After parsing - Lines: %t, Bytes: %t\n", *countLines, *countBytes)

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
	fmt.Printf("Write your text 'exit' to finish (Counting %s):", mode)

	// read user input
	input := readInput()

	// call the counting function
	total := countString(input, *countLines, *countBytes)

	switch mode {
	case "lines":
		fmt.Printf("Lines: %d\n ", total)
	case "bytes":
		fmt.Printf("bytes: %d\n ", total)
	default:
		fmt.Printf("Words: %d\n ", total)
	}

}
