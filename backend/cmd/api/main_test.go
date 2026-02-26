package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckRoute(t *testing.T) {
	// 1. Setup our router (we can pass nil for the handler since the health route doesn't use it)
	router := setupRouter(nil)

	// 2. Create a response recorder (to "catch" the response)
	w := httptest.NewRecorder()

	// 3. Create a mock request
	req, _ := http.NewRequest("GET", "/health", nil)

	// 4. Perform the request
	router.ServeHTTP(w, req)

	// 5. Assertions using Testify
	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200")

	// Parse the JSON response
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err, "Response should be valid JSON")
	assert.Equal(t, "up", response["status"])
	assert.Equal(t, "Movie Streaming API with Gin", response["message"])
}
