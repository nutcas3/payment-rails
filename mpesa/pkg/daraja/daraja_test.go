package daraja

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	service, err := New("test-api-key", "test-consumer-secret", "test-pass-key", SANDBOX)
	if err != nil {
		t.Fatalf("Failed to create service with valid parameters: %v", err)
	}
	if service == nil {
		t.Fatal("Service should not be nil")
	}
	if service.apiKey != "test-api-key" {
		t.Errorf("Expected apiKey to be 'test-api-key', got '%s'", service.apiKey)
	}
	if service.consumerSecret != "test-consumer-secret" {
		t.Errorf("Expected consumerSecret to be 'test-consumer-secret', got '%s'", service.consumerSecret)
	}
	if service.passKey != "test-pass-key" {
		t.Errorf("Expected passKey to be 'test-pass-key', got '%s'", service.passKey)
	}
	if service.environment != SANDBOX {
		t.Errorf("Expected environment to be SANDBOX, got '%s'", service.environment)
	}
	if service.baseURL != "https://sandbox.safaricom.co.ke" {
		t.Errorf("Expected baseURL to be 'https://sandbox.safaricom.co.ke', got '%s'", service.baseURL)
	}
	if service.httpClient == nil {
		t.Error("HTTP client should not be nil")
	}
	if service.cache == nil {
		t.Error("Cache should not be nil")
	}

	service, err = New("", "test-consumer-secret", "test-pass-key", SANDBOX)
	if err == nil {
		t.Error("Expected error for empty apiKey, got nil")
	}
	if service != nil {
		t.Error("Service should be nil for invalid parameters")
	}

	service, err = New("test-api-key", "", "test-pass-key", SANDBOX)
	if err == nil {
		t.Error("Expected error for empty consumerSecret, got nil")
	}
	if service != nil {
		t.Error("Service should be nil for invalid parameters")
	}

	service, err = New("test-api-key", "test-consumer-secret", "", SANDBOX)
	if err == nil {
		t.Error("Expected error for empty passKey, got nil")
	}
	if service != nil {
		t.Error("Service should be nil for invalid parameters")
	}

	service, err = New("test-api-key", "test-consumer-secret", "test-pass-key", PRODUCTION)
	if err != nil {
		t.Fatalf("Failed to create service with production environment: %v", err)
	}
	if service.baseURL != "https://api.safaricom.co.ke" {
		t.Errorf("Expected baseURL to be 'https://api.safaricom.co.ke', got '%s'", service.baseURL)
	}
}

func TestSetHttpClient(t *testing.T) {
	service, _ := New("test-api-key", "test-consumer-secret", "test-pass-key", SANDBOX)
	
	customClient := &http.Client{
		Timeout: 60 * time.Second,
	}
	
	service.SetHttpClient(customClient)
	
	if service.httpClient != customClient {
		t.Error("SetHttpClient did not set the custom HTTP client")
	}
}

func TestGetAuthToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/v1/generate" {
			t.Errorf("Expected request to '/oauth/v1/generate', got '%s'", r.URL.Path)
		}
		
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			t.Error("Authorization header is missing")
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token":"test-access-token","expires_in":"3599"}`))
	}))
	defer server.Close()
	
	service, _ := New("test-api-key", "test-consumer-secret", "test-pass-key", SANDBOX)
	service.baseURL = server.URL
	
	token, err := service.GetAuthToken()
	if err != nil {
		t.Fatalf("GetAuthToken failed: %v", err)
	}
	
	if token != "test-access-token" {
		t.Errorf("Expected token to be 'test-access-token', got '%s'", token)
	}
	
	// Test cached token
	server.Close() // Close server to ensure we're using the cached token
	token, err = service.GetAuthToken()
	if err != nil {
		t.Fatalf("GetAuthToken from cache failed: %v", err)
	}
	
	if token != "test-access-token" {
		t.Errorf("Expected cached token to be 'test-access-token', got '%s'", token)
	}
}

func TestMakeRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth/v1/generate" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"access_token":"test-access-token","expires_in":"3599"}`))
			return
		}
		
		if r.URL.Path == "/test/endpoint" {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "Bearer test-access-token" {
				t.Errorf("Expected Authorization header 'Bearer test-access-token', got '%s'", authHeader)
			}
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result":"success"}`))
			return
		}
		
		t.Errorf("Unexpected request to '%s'", r.URL.Path)
	}))
	defer server.Close()
	
	service, _ := New("test-api-key", "test-consumer-secret", "test-pass-key", SANDBOX)
	service.baseURL = server.URL
	
	resp, err := service.makeRequest(http.MethodGet, "/test/endpoint", nil)
	if err != nil {
		t.Fatalf("makeRequest failed: %v", err)
	}
	
	expected := `{"result":"success"}`
	if string(resp) != expected {
		t.Errorf("Expected response '%s', got '%s'", expected, string(resp))
	}
}
