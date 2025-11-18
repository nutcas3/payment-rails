package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticate_FormEncoded(t *testing.T) {
	// Mock server to verify request format
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/token" {
			t.Errorf("Expected path /oauth/token, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf("Expected Content-Type application/x-www-form-urlencoded, got %s", r.Header.Get("Content-Type"))
		}

		err := r.ParseForm()
		if err != nil {
			t.Fatalf("Failed to parse form: %v", err)
		}

		if r.Form.Get("grant_type") != "client_credentials" {
			t.Errorf("Expected grant_type client_credentials, got %s", r.Form.Get("grant_type"))
		}
		if r.Form.Get("client_id") != "test-client" {
			t.Errorf("Expected client_id test-client, got %s", r.Form.Get("client_id"))
		}
		if r.Form.Get("client_secret") != "test-secret" {
			t.Errorf("Expected client_secret test-secret, got %s", r.Form.Get("client_secret"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "test-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	client := NewClient(&ClientConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		BaseURL:      server.URL,
	})

	err := client.Authenticate(context.Background())
	if err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}
}

func TestTokenExpiry_ShortLived(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "short-lived-token",
			"token_type":   "Bearer",
			"expires_in":   30,
		})
	}))
	defer server.Close()

	client := NewClient(&ClientConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		BaseURL:      server.URL,
	})

	err := client.Authenticate(context.Background())
	if err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}

	if !client.isTokenValid() {
		t.Error("Token should be valid immediately after auth")
	}
}
