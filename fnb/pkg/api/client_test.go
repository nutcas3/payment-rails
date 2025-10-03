package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	config := &ClientConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		APIKey:       "test-api-key",
		Environment:  "sandbox",
	}

	client := NewClient(config)

	if client == nil {
		t.Fatal("Expected client to be created")
	}

	if client.clientID != config.ClientID {
		t.Errorf("Expected clientID %s, got %s", config.ClientID, client.clientID)
	}

	if client.baseURL != SandboxBaseURL {
		t.Errorf("Expected baseURL %s, got %s", SandboxBaseURL, client.baseURL)
	}
}

func TestAuthenticate(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/token" {
			t.Errorf("Expected path /oauth/token, got %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		// Return mock auth response
		resp := AuthResponse{
			AccessToken: "mock-access-token",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	config := &ClientConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		APIKey:       "test-api-key",
		BaseURL:      server.URL,
	}

	client := NewClient(config)
	ctx := context.Background()

	err := client.Authenticate(ctx)
	if err != nil {
		t.Fatalf("Authentication failed: %v", err)
	}

	if client.accessToken != "mock-access-token" {
		t.Errorf("Expected access token 'mock-access-token', got %s", client.accessToken)
	}

	if !client.isTokenValid() {
		t.Error("Expected token to be valid")
	}
}

func TestAuthenticateError(t *testing.T) {
	// Create mock server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		resp := ErrorResponse{
			ErrorCode:        "invalid_client",
			ErrorDescription: "Invalid client credentials",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	config := &ClientConfig{
		ClientID:     "invalid-client-id",
		ClientSecret: "invalid-secret",
		APIKey:       "test-api-key",
		BaseURL:      server.URL,
	}

	client := NewClient(config)
	ctx := context.Background()

	err := client.Authenticate(ctx)
	if err == nil {
		t.Fatal("Expected authentication to fail")
	}

	errResp, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("Expected ErrorResponse, got %T", err)
	}

	if errResp.ErrorCode != "invalid_client" {
		t.Errorf("Expected error 'invalid_client', got %s", errResp.ErrorCode)
	}
}

func TestDoRequest(t *testing.T) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth/token" {
			// Auth endpoint
			resp := AuthResponse{
				AccessToken: "mock-access-token",
				TokenType:   "Bearer",
				ExpiresIn:   3600,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}

		if r.URL.Path == "/api/v1/test" {
			// Test endpoint
			if r.Header.Get("Authorization") != "Bearer mock-access-token" {
				t.Error("Expected Authorization header with bearer token")
			}

			if r.Header.Get("X-API-Key") == "" {
				t.Error("Expected X-API-Key header")
			}

			resp := map[string]string{
				"status": "success",
				"data":   "test-data",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	config := &ClientConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		APIKey:       "test-api-key",
		BaseURL:      server.URL,
	}

	client := NewClient(config)
	ctx := context.Background()

	var result map[string]string
	err := client.DoRequest(ctx, "GET", "/api/v1/test", nil, &result)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if result["status"] != "success" {
		t.Errorf("Expected status 'success', got %s", result["status"])
	}

	if result["data"] != "test-data" {
		t.Errorf("Expected data 'test-data', got %s", result["data"])
	}
}

func TestTokenExpiry(t *testing.T) {
	config := &ClientConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		APIKey:       "test-api-key",
		Environment:  "sandbox",
	}

	client := NewClient(config)

	// Set expired token
	client.accessToken = "expired-token"
	client.tokenExpiry = time.Now().Add(-1 * time.Hour)

	if client.isTokenValid() {
		t.Error("Expected token to be invalid")
	}

	// Set valid token
	client.accessToken = "valid-token"
	client.tokenExpiry = time.Now().Add(1 * time.Hour)

	if !client.isTokenValid() {
		t.Error("Expected token to be valid")
	}
}

func TestErrorResponse(t *testing.T) {
	err := &ErrorResponse{
		ErrorCode:        "invalid_request",
		ErrorDescription: "Missing required parameter",
		Status:           400,
	}

	expected := "FNB API error: invalid_request - Missing required parameter"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}
