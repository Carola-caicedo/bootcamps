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
			name:     "When the user types a single sentence with some words./Cuando escribe una sola oración con algunas palabras.",
			input:    "Hola mundo, es un ejemplo con un simple texto",
			expected: 9,
		},

		{
			name: "When the user types multiple sentences with some words per sentence./Cuando el usuario escribe varias oraciones con algunas palabras por oración.",
			input: `Esta es una oración con algunas palabras.
				Esta es otra oración con más palabras.
				Finalmente, esta es la última oración.`,
			expected: 20,
		},

		{
			name:     "When the user types a single word./Cuando el usuario escribe solo una palabra.",
			input:    "hello",
			expected: 1,
		},

		{
			name:     "When the user types a single composed word./Cuando el usuario escribe una sola palabra compuesta.",
			input:    "read-only",
			expected: 1,
		},

		{
			name:     "When the user types multiple break lines./Cuando el usuario escribe varias líneas de salto.",
			input:    "This is a line of text.\nStart a line break here.\nesta es una linea de texto.\naqui da un salto de linea.\n",
			expected: 23,
		},

		{
			name:     "When te user types (Exit), (exit) or (EXIT)./ Cuando el usuario usa tipos como (Exit), (exit) or (EXIT).",
			input:    "I'm going to use the commands that the program uses for output, Exit, exit, EXIT./ Voy a usar los comandos que el programa utiliza para salir.",
			expected: 26,
		},

		{
			name:     "Empty input./Entrada vacía.",
			input:    "",
			expected: 0,
		},
	}

	//that yields the expected result
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countString(tt.input, false)
			if result != tt.expected {
				t.Errorf("For input '%s': expected %d, got %d", tt.input, tt.expected, result)
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
			name:     "Single Line./Una sola línea.",
			input:    "Hello world, this is a single line of text./Hola mundo, es un ejemplo con un simple texto",
			expected: 1,
		},

		{
			name:     "Multiple Lines./Varias líneas.",
			input:    "This is a line of text.\nStart a line break here.\nesta es una linea de texto.\naqui da un salto de linea.\n",
			expected: 4,
		},

		{
			name:     "Multiple lines with breaks./Varias líneas con saltos de línea.",
			input:    "This is a line of text.\n\nStart a line break here.\n\nEsta es una linea de texto.\n\naqui da un salto de linea.\n\n",
			expected: 8,
		},

		{
			name:     "Empty input./Entrada vacía.",
			input:    "",
			expected: 0,
		},
	}

	//that yields the expected result
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countString(tt.input, true)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: expected '%d', got '%d'", tt.input, tt.expected, result)
			}
		})
	}
}

//additional function, which accepts all types of outputs written in the program
func TestCountString_Exit(t *testing.T) {
	//What is the input/data to be tested, minus data (name)
	test := []struct {
		input    string
		expected int
	}{
		{"exit", 0},
		{"Exit", 0},
		{"EXIT", 0},
	}

	//What is expected in the test outcome/result
	for _, tt := range test {
		t.Run(tt.input, func(t *testing.T) {
			result := countString(tt.input, false)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: expected '%d', got '%d'", tt.input, tt.expected, result)
			}

		})
	}

}
