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

type testCase struct {
	name            string
	path            string
	expectedCode    int
	expectedContent string
}

func TestGet(t *testing.T) {
	resultCases := []testCase{
		{
			name:            "Root returns Error 404",
			path:            "/",
			expectedCode:    http.StatusNotFound,
			expectedContent: "404 page not found",
		},

		{
			name:            "Root returns 200 OK",
			path:            "/anything",
			expectedCode:    http.StatusOK,
			expectedContent: "Hello World",
		},
	}

	url, cleaner := setupAPI(t)
	defer cleaner()

	for _, tc := range resultCases {
		t.Run(tc.name, func(t *testing.T) {

			r, err := http.Get(url + tc.path)
			if err != nil {
				t.Fatal(err)
			}

			defer r.Body.Close()

			if r.StatusCode != tc.expectedCode {
				t.Errorf("Expected %d , got %d", tc.expectedCode, r.StatusCode)
			}

			ContentType := r.Header.Get("Content-Type")
			switch ContentType {
			case "text/plain; charset=utf-8":
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}

				if !strings.Contains(string(body), tc.expectedContent) {
					t.Errorf("Expected %s, got %s", tc.expectedContent, string(body))
				}

			default:
				t.Fatalf("Unsupported Content-Type: %q", r.Header.Get("Content-Type"))

			}

		})
	}

}
