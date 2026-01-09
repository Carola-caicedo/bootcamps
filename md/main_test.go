package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {

	// Test with a valid filename
	t.Run("ValidFile", func(t *testing.T) {
		// temporary file
		testFile := "test.html"

		err := run(testFile)
		if err != nil {
			t.Fatalf("run() failed: %v", err)
		}

		// Check if the file created
		expectedFilename := testFile + ".html"
		if _, err := os.Stat(expectedFilename); os.IsNotExist(err) {
			t.Errorf("The file %s doesn't exits", expectedFilename)
		}
		// Clean
		os.Remove(expectedFilename)
	})

	// Test with an empty filename
	t.Run("EmptyFile", func(t *testing.T) {
		err := run("")
		if err == nil {
			t.Error("Expected an error for empty filename")
		}

		expectedError := "the flag (-out) is obligatory "
		if err != nil && err.Error() != expectedError {
			t.Errorf("Failed: expected %q, got %q", expectedError, err)
		}

	})
}
