package main

import "testing"

// function to count words with different conditions
func TestCountString_Words(t *testing.T) {

	//What is the input/data to be tested
	tests := []struct {
		name     string
		input    string
		expected int
	}{

		{
			name:     "single sentence with words.",
			input:    "A single sentence with a few words.",
			expected: 7,
		},

		{
			name: "multiple sentences with some words ",
			input: `This is a line of text.
				Start a line break here.`,
			expected: 11, //11?
		},

		{
			name:     "single word.",
			input:    "hello",
			expected: 1,
		},

		{
			name:     "single composed word.",
			input:    "read-only",
			expected: 1,
		},

		{
			name:     "multiple break lines.",
			input:    "This is a line of text.\n\nStart a line break here.",
			expected: 11, //11?
		},

		{
			name:     "types (Exit), (exit) or (EXIT).",
			input:    "that the program uses for output, Exit, exit, EXIT.",
			expected: 9,
		},

		{
			name:     "Empty input.",
			input:    "",
			expected: 0,
		},
	}

	//that yields the expected result
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countString(tt.input, false)
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
		{
			name:     "Single Line.",
			input:    "Hello world, this is a single line of text.",
			expected: 1,
		},

		{
			name:     "Multiple Lines.",
			input:    "This is a line of text.\nStart a line break here.",
			expected: 2,
		},

		{
			name:     "Multiple lines with breaks.",
			input:    "This is a line of text.\n\nStart a line break here.",
			expected: 3,
		},

		{
			name:     "Empty input",
			input:    "",
			expected: 0,
		},
	}

	//that yields the expected result
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countString(tt.input, true)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: expected '%d', got '%d'", tt.name, tt.expected, result)
			}
		})
	}
}

//additional function, which accepts all types of outputs written in the program
func TestCountString_Exit(t *testing.T) {
	//What is the input/data to be tested, minus data (name)
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
			result := countString(tt.input, false)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: expected '%d', got '%d'", tt.name, tt.expected, result)
			}

		})
	}

}
