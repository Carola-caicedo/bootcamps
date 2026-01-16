package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun(t *testing.T) {

	t.Run("ErrWithoutIn", func(t *testing.T) {
		err := run("", "")
		if err == nil {
			t.Error("expected error, when -in is not specified")
		}
	})
	// file created (-out)
	t.Run("FileCreated", func(t *testing.T) {
		tmpFile := "testfile.md"
		err := os.WriteFile(tmpFile, []byte("# test"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(tmpFile)

		err = run(tmpFile, "result")
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

	t.Run("FileCreatedWithoutOut", func(t *testing.T) {
		tmpFile := "document.md"
		err := os.WriteFile(tmpFile, []byte("# test"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(tmpFile)

		err = run(tmpFile, "")
		if err != nil {
			t.Fatalf("run() failed: %v", err)
		}

		if _, err := os.Stat("document.html"); os.IsNotExist(err) {
			t.Errorf("The file document.html was not created")
		} else {
			// Clean
			os.Remove("document.html")
		}
	})
}

func TestParseContent(t *testing.T) {

	mdContent := []byte(`# Test Markdown File

Just a test

## Bullets

- Links [Link1](https://example.com)

## Quotes

> Quotes in **bold** and _italic_ text`)

	mdFile := "test_parse.md"
	err := os.WriteFile(mdFile, mdContent, 0644)
	if err != nil {
		t.Fatalf("failed create test file: %v", err)
	}
	defer os.Remove(mdFile)

	outputBaseName := "test_parse"
	err = run(mdFile, outputBaseName)
	if err != nil {
		t.Fatalf("run() failed: %v", err)
	}

	outputFile := outputBaseName + ".html"
	defer os.Remove(outputFile)

	generatedHTML, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed read generated file: %v", err)
	}

	goldenHTML, err := os.ReadFile("test_golden.html")
	if err != nil {
		t.Fatalf("failed read golden file: %v", err)
	}

	normalize := func(data []byte) []byte {
		return bytes.TrimSpace(data)
	}

	// Comparar
	if !bytes.Equal(normalize(generatedHTML), normalize(goldenHTML)) {
		t.Error("Generated HTML does not match golden file")

		// Opcional: imprimir ambos para debugging
		t.Logf("Generated (%d bytes):\n%s", len(generatedHTML), generatedHTML)
		t.Logf("Golden (%d bytes):\n%s", len(goldenHTML), goldenHTML)

		// También mostrar las longitudes de los normalizados
		t.Logf("Normalized generated: %d bytes", len(normalize(generatedHTML)))
		t.Logf("Normalized golden: %d bytes", len(normalize(goldenHTML)))
	}

}
