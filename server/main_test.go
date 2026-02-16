package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()

	server := httptest.NewServer(newMux())

	url = server.URL

	cleaner = func() {
		server.Close()
	}

	return url, cleaner
}

func TestRootHandler(t *testing.T) {

	path := "/"
	expectedCode := http.StatusNotFound
	expectedContent := "404 page not found"

	url, cleaner := setupAPI(t)
	defer cleaner()

	resp, err := http.Get(url + path)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != expectedCode {
		t.Errorf("Expected %d , got %d", expectedCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), expectedContent) {
		t.Errorf("Expected %s, got %s", expectedContent, string(body))
	}

}

func TestGetNotFound(t *testing.T) {
	path := "/pathnotexist"
	expectedCode := http.StatusOK
	expectedContent := "Hello World"

	url, cleaner := setupAPI(t)
	defer cleaner()

	resp, err := http.Get(url + path)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != expectedCode {
		t.Errorf("Expected %d, got %d", expectedCode, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), expectedContent) {
		t.Errorf("Expected %s, got %s", expectedContent, string(body))
	}

}
