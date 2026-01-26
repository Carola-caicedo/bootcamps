package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunOut(t *testing.T) {
	tmpDir := t.TempDir()
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(tmpDir)

	if err := os.MkdirAll("md", 0755); err != nil {
		t.Fatal(err)
	}

	mdFile := "testfile.md"
	mdContent := []byte("# This is a test.")
	if err := os.WriteFile(mdFile, mdContent, 0644); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	err := run(mdFile, "result", &buf)
	if err != nil {
		t.Fatalf("run() failed: %v", err)
	}

	expectedResult := filepath.Join("md", "result.html\n")
	if got := buf.String(); got != expectedResult {
		t.Errorf("expected %q, got %q", expectedResult, got)
	}

	if _, err := os.Stat(filepath.Join("md", "result.html")); os.IsNotExist(err) {
		t.Errorf("The file result.html was not created")
	}
}

func TestRunWithoutOut(t *testing.T) {
	tmpDir := t.TempDir()
	oldDir, _ := os.Getwd()

	defer os.Chdir(oldDir)
	os.Chdir(tmpDir)

	if err := os.MkdirAll("md", 0755); err != nil {
		t.Fatal(err)
	}

	mdFile := "testfile.md"
	mdContent := []byte("# This is a test.")
	if err := os.WriteFile(mdFile, mdContent, 0644); err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := run(mdFile, "", &buf); err != nil {
		t.Fatalf("run() failed: %v", err)
	}

	filename := strings.TrimSpace(buf.String())
	if filename == "" {
		t.Errorf("No filename")
	}

	//For checking the filename: method CONTAINS
	if !strings.HasPrefix(filepath.Base(filename), "md") || !strings.HasSuffix(filename, ".html") {
		t.Errorf("Not matching filename: %q", filename)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("The file %q was not created", filename)
	}
}

func TestParseContent(t *testing.T) {

	goldenpath := filepath.Join("testdata", "test_golden.html")

	mdContent := []byte(`# Test Markdown File

Just a test

## Bullets

- Links [Link1](https://example.com)

## Quotes

> Quotes in **bold** and _italic_ text`)

	tmpDir := t.TempDir()
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(tmpDir)

	if err := os.MkdirAll("md", 0755); err != nil {
		t.Fatal(err)
	}

	mdFile := "test_parse.md"
	err := os.WriteFile(mdFile, mdContent, 0644)
	if err != nil {
		t.Fatalf("failed create test file: %v", err)
	}
	defer os.Remove(mdFile)

	var buf bytes.Buffer
	outputBaseName := "test_parse"
	err = run(mdFile, outputBaseName, &buf)
	if err != nil {
		t.Fatalf("run() failed: %v", err)
	}

	outputFile := filepath.Join("md", outputBaseName+".html")
	defer os.Remove(outputFile)

	generatedHTML, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed read generated file: %v", err)
	}

	os.Chdir(oldDir)

	goldenHTML, err := os.ReadFile(goldenpath)
	if err != nil {
		t.Fatalf("failed read golden file: %v", err)
	}

	normalize := func(data []byte) []byte {
		return bytes.TrimSpace(data)
	}

	// Compare
	if !bytes.Equal(normalize(generatedHTML), normalize(goldenHTML)) {
		t.Error("Generated HTML does not match golden file")

		// print both for debugging
		t.Logf("Generated (%d bytes):\n%s", len(generatedHTML), generatedHTML)
		t.Logf("Golden (%d bytes):\n%s", len(goldenHTML), goldenHTML)

		// Show the lengths
		t.Logf("Normalized generated: %d bytes", len(normalize(generatedHTML)))
		t.Logf("Normalized golden: %d bytes", len(normalize(goldenHTML)))
	}

}
