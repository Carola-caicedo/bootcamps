package main

import (
	"os"
	"testing"
	"bytes"
)

func TestRun(t *testing.T) {

	t.Run("ErrWithoutIn", func(t *testing.T) {
		err := run("", "")
		if err == nil {
			t.Error("expected error, when -in is not specified")
		}
	})
		// file created
		t.Run ("FileCreated", func(t *testing.T) {
			os.WriteFile("README.md", []byte("# test"), 0644)
			defer os.Remove("README.md")

			err := run("README.md", "result")
			if err != nil {
				t.Fatalf("run() failed: %v", err)
			}


			if _, err := os.Stat("result.html"); os.IsNotExist(err) {
				t.Errorf("The file result.html was not created")
			} else {
			// Clean
			os.Remove("result.html")
			}
		})
}

func TestParseContent(t *testing.T){
	mdBytes, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	result, err := parseContent(mdBytes)
	if err != nil {
		t.Fatalf("Error parsing content: %v", err)
	}

	goldenBytes, err := os.ReadFile("test_golden.html")
	if err != nil {
		t.Fatalf("Error reading golden file: %v", err)
	}

	if !bytes.Equal(result, goldenBytes){
		t.Error("The result does not match withthe golden file")
	}
}
