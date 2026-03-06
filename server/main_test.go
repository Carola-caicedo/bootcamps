package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type todoResponseJSON struct {
	Results      []todoItem `json:"results"`
	Date         string     `json:"date"`
	TotalResults int        `json:"total_results"`
}

type todoItem struct {
	Task        string `json:"task"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at"`
}

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "testdata*.json")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}

	initData := []byte(`[]`)
	if _, err := tempFile.Write(initData); err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}

	tempFile.Close()

	server := httptest.NewServer(newMux(tempFile.Name()))

	for i := 0; i < 3; i++ {
		taskName := fmt.Sprintf("Task %d", i)

		taskJSON := fmt.Sprintf(`{"task":"%s"}`, taskName)

		resp, err := http.Post(server.URL+"/todo", "application/json", bytes.NewBufferString(taskJSON))

		if err != nil {
			t.Fatalf("Error posting task %d: %v", i, err)
		}

		if resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 201 Created task %d, got %d, body: %s", i, resp.StatusCode, string(body))
		}

		resp.Body.Close()
	}

	url = server.URL

	cleaner = func() {
		server.Close()
		os.Remove(tempFile.Name())
	}

	return url, cleaner
}

type testCase struct {
	name            string
	path            string
	expectedCode    int
	expectedContent string
	expectItems     int
}

func TestGet(t *testing.T) {
	resultCases := []testCase{
		{
			name:            "Root returns Error 404",
			path:            "/",
			expectedCode:    http.StatusNotFound,
			expectedContent: "404 page not found",
			expectItems:     0,
		},

		{
			name:            "Root returns 200 OK",
			path:            "/anything",
			expectedCode:    http.StatusOK,
			expectedContent: "Hello World",
			expectItems:     0,
		},

		{
			name:            "Get all returns 200 OK",
			path:            "/todo",
			expectedCode:    http.StatusOK,
			expectedContent: "",
			expectItems:     3,
		},

		{
			name:            "Get one",
			path:            "/todo/1",
			expectedCode:    http.StatusOK,
			expectedContent: "",
			expectItems:     1,
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
			switch {
			case strings.Contains(ContentType, "text/plain; charset=utf-8"):
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}

				if !strings.Contains(string(body), tc.expectedContent) {
					t.Errorf("Expected %s, got %s", tc.expectedContent, string(body))
				}

			case strings.Contains(ContentType, "application/json"):
				var response todoResponseJSON
				err := json.NewDecoder(r.Body).Decode(&response)
				if err != nil {
					t.Fatalf("Error decoding JSON: %v", err)
				}

				if response.TotalResults != tc.expectItems {
					t.Errorf("Expected %d items, got %d", tc.expectItems, response.TotalResults)
				}

				if len(response.Results) != tc.expectItems {
					t.Errorf("Expected %d items, got %d", tc.expectItems, len(response.Results))
				}

				if tc.name == "Get one" && len(response.Results) > 0 {
					expectedTask := "Task 1"
					if response.Results[0].Task != expectedTask {
						t.Errorf("Expected task %s, got %s", expectedTask, response.Results[0].Task)
					}
				}

				if tc.name == "Get all returns 200 OK" && len(response.Results) > 0 {
					expectedTask := "Task 0"
					if response.Results[0].Task != expectedTask {
						t.Errorf("Expected task %s, got %s", expectedTask, response.Results[0].Task)
					}
				}

			default:
				t.Fatalf("Unsupported Content-Type: %q", r.Header.Get("Content-Type"))

			}

		})
	}

}

func TestAdd(t *testing.T) {

	url, cleaner := setupAPI(t)
	defer cleaner()

	t.Run("Add new task", func(t *testing.T) {
		newTaskJSON := `{"task":"Task 3"}`

		resp, err := http.Post(url+"/todo", "application/json", bytes.NewBufferString(newTaskJSON))

		if err != nil {
			t.Fatalf("Error posting task: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 201 Created, got %d, body: %s", resp.StatusCode, string(body))
		}
	})

	t.Run("CheckAdd", func(t *testing.T) {
		resp, err := http.Get(url + "/todo/3")
		if err != nil {
			t.Fatalf("Error getting task: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 200 OK, got %d, body: %s", resp.StatusCode, string(body))
		}

		contentType := resp.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			t.Fatalf("Expected Content-Type application/json, got %s", contentType)
		}

		var response todoResponseJSON
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Error decoding JSON: %v", err)
		}

		if len(response.Results) != 1 {
			t.Fatalf("Expected 1 item, got %d", len(response.Results))
		}

		if response.TotalResults != 1 {
			t.Fatalf("Expected 1 item, got %d", response.TotalResults)
		}

		expectedTask := "Task 3"
		if response.Results[0].Task != expectedTask {
			t.Fatalf("Expected task %s, got %s", expectedTask, response.Results[0].Task)
		}
	})
}

func TestDelete(t *testing.T) {
	url, cleaner := setupAPI(t)
	defer cleaner()

	t.Run("Delete", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, url+"/todo/0", nil)
		if err != nil {
			t.Fatalf("Error deleting task: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error deleting task: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 204 No Content, got %d, body: %s", resp.StatusCode, string(body))
		}
	})

	t.Run("CheckDelete", func(t *testing.T) {
		resp, err := http.Get(url + "/todo")
		if err != nil {
			t.Fatalf("Error getting task: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 200 ok, got %d, body: %s", resp.StatusCode, string(body))
		}

		contentType := resp.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			t.Fatalf("Expected Content-Type application/json, got %s", contentType)
		}

		var response todoResponseJSON
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Error decoding JSON: %v", err)
		}

		expectedItems := 2
		if response.TotalResults != expectedItems {
			t.Fatalf("Expected %d items, got %d", expectedItems, response.TotalResults)
		}

		if len(response.Results) != expectedItems {
			t.Fatalf("Expected %d items, got %d", expectedItems, len(response.Results))
		}

		if len(response.Results) > 0 {
			expectedTask := "Task 1"
			if response.Results[0].Task != expectedTask {
				t.Fatalf("Expected task %s, got %s", expectedTask, response.Results[0].Task)
			}
		}
	})
}

func TestComplete(t *testing.T) {
	url, cleaner := setupAPI(t)
	defer cleaner()

	t.Run("Complete", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPatch, url+"/todo/0?complete", nil)
		if err != nil {
			t.Fatalf("Error created task: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error executing patch: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNoContent {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 204 No Content, got %d, body: %s", resp.StatusCode, string(body))
		}
	})

	t.Run("CheckComplete", func(t *testing.T) {
		resp, err := http.Get(url + "/todo")
		if err != nil {
			t.Fatalf("Error getting task: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 200 ok, got %d, body: %s", resp.StatusCode, string(body))
		}

		contentType := resp.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			t.Fatalf("Expected Content-Type application/json, got %s", contentType)
		}

		var response todoResponseJSON
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Error decoding JSON: %v", err)
		}

		if response.TotalResults != 3 {
			t.Fatalf("Expected 3 items, got %d", response.TotalResults)
		}

		if len(response.Results) != 3 {
			t.Fatalf("Expected 3 items, got %d", len(response.Results))
		}

		if len(response.Results) > 0 {
			if !response.Results[0].Done {
				t.Fatalf("Expected a single task Done=true, got %v", response.Results[0].Done)
			}

			if response.Results[0].CompletedAt == "0001-01-01T00:00:00Z" {
				t.Fatalf("Expected a single task Completed, got zero")
			}
		}

		for i := 1; i < len(response.Results); i++ {
			if response.Results[i].Done {
				t.Fatalf("Expected task %d done=false, got %v", i, response.Results[i].Done)
			}

			if response.Results[i].CompletedAt != "0001-01-01T00:00:00Z" {
				t.Fatalf("Expected task %d Completed at zero, got %q", i, response.Results[i].CompletedAt)
			}
		}
	})

}
