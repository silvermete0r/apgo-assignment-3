package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test User Registration API
func TestRegisterUser(t *testing.T) {
	app := &application{}

	user := map[string]string{
		"name":     "John Doe",
		"email":    "johndoe@example.com",
		"password": "password123",
	}
	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	app.routes().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status 201 Created")
}

// Define a custom testServer type which anonymously embeds a httptest.Server instance.
func TestAuthenticateUser(t *testing.T) {
	app := &application{}

	user := map[string]string{
		"email":    "johndoe@example.com",
		"password": "password123",
	}
	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/v1/tokens/authentication", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	app.routes().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status 200 OK")
}

// Test Protected Endpoint Access
func TestProtectedEndpoint(t *testing.T) {
	app := &application{}

	req, _ := http.NewRequest(http.MethodGet, "/v1/protected", nil)
	req.Header.Set("Authorization", "Bearer valid.jwt.token")
	rr := httptest.NewRecorder()

	app.routes().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status 200 OK")
}
