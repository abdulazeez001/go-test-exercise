package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	HelloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", rec.Code)
	}

	expected := "Hello, World!\n"
	if rec.Body.String() != expected {
		t.Errorf("Expected body %v, got %v", expected, rec.Body.String())
	}
}

func TestHelloHandlerWithDifferentMethods(t *testing.T) {
	methods := []string{"POST", "PUT", "DELETE", "PATCH"}

	for _, method := range methods {
		req, err := http.NewRequest(method, "/", nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		helloHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status OK for method %v, got %v", method, rec.Code)
		}

		expected := "Hello, World!\n"
		if rec.Body.String() != expected {
			t.Errorf("Expected body %v for method %v, got %v", expected, method, rec.Body.String())
		}
	}
}

func TestMainHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	methods := []string{"POST", "PUT", "DELETE", "PATCH"}

	for _, method := range methods {
		req, err := http.NewRequest(method, ts.URL, nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Could not send %v request: %v", method, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status OK for method %v, got %v", method, resp.StatusCode)
		}
	}
}
