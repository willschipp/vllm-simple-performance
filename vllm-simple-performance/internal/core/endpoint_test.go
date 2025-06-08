package core

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestSendPrompt(t *testing.T) {
	// Set up a test HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check method and content type
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected application/json, got %s", r.Header.Get("Content-Type"))
		}
		// Read and check the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}
		var payload Payload
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("Failed to unmarshal request body: %v", err)
		}
		if payload.Model != "test-model" || payload.Prompt != "test-prompt" {
			t.Errorf("Unexpected payload: %+v", payload)
		}
		// Write a dummy JSON response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result":"success"}`))
	}))
	defer ts.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	SendPrompt(ts.URL, "test-model", "test-prompt", &wg)
	wg.Wait()
}