package main

import (
	"testing"
)

// function to count words with different conditions
func TestCountString_Words(t *testing.T) {

	//What is the input/data to be tested
	tests := []struct {
		name     string
		input    string
		expected int
	}{

		{"TestSingleSentence", "A single sentence with a few words.", 7},
		{"TestMultipleSentences", "This is a line of text.\n\t\t\tStart a line break here.", 11},
		{"TestSingleWord", "hello", 1},
		{"TestSingleComposedWord", "read-only", 1},
		{"TestMultipleBreakLines", "This is a line of text.\n\nStart a line break here.", 11},
		{"TestExitVarations", "that the program uses for output, Exit, exit, EXIT.", 9},
		{"TestEmptyInput", "", 0},
	}

	//the expected result
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countString(tt.input, false, false)
			if result != tt.expected {
				t.Errorf("For input '%s': expected %d, got %d", tt.name, tt.expected, result)
			}
		})

	}
}

// function to count lines with different conditions
func TestCountString_Lines(t *testing.T) {
	//What is the input/data to be tested
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"TestSingleLine", "Hello world, this is a single line of text.", 1},
		{"TestMultipleLines", "This is a line of text.\nStart a line break here.", 2},
		{"TestMultipleLinesWithBreaks", "This is a line of text.\n\nStart a line break here.", 3},
		{"TestEmptyInput", "", 0},
	}

	//the expected result
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countString(tt.input, true, false)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: expected '%d', got '%d'", tt.name, tt.expected, result)
			}
		})
	}
}

// function to count bytes with different conditions
func TestCountString_Bytes(t *testing.T) {
	//What is the input/data to be tested
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"TestSingleLineBytes", "Hello world, this is a single line of text.", 43},
		{"TestMultipleLinesBytes", "This is a line of text.\nStart a line break here.", 48},
		{"TestMultipleLinesWithBreaksBytes", "This is a line of text.\n\nStart a line break here.", 49},
		{"TestEmptyInputBytes", "", 0},
	}

	//the expected result
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countString(tt.input, false, true)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: expected '%d', got '%d'", tt.name, tt.expected, result)
			}
		})
	}
}

// Test aditional verufy priority countLines > countBytes > countWords
func TestCountString_Priority(t *testing.T) {
	//What is the input/data to be tested
	input := "Line1\nLine2\nLine3\n"

	//When same flags are true, countLines should take priority
	result := countString(input, true, true)
	expected := 3

	if result != expected {
		t.Errorf("Test 'Priority' failed: expected '%d', got '%d'", expected, result)
	}
}

// additional function, which accepts all types of outputs written in the program
func TestCountString_Exit(t *testing.T) {
	//What is the input/data to be tested, name data
	test := []struct {
		name     string
		input    string
		expected int
	}{
		{"exit", "exit", 0},
		{"Exit", "Exit", 0},
		{"EXIT", "EXIT", 0},
	}

	//What is expected in the test outcome/result
	for _, tt := range test {
		t.Run(tt.input, func(t *testing.T) {
			result := countString(tt.input, false, false)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: expected '%d', got '%d'", tt.name, tt.expected, result)
			}

		})
	}

}
