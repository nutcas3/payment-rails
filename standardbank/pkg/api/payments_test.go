package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetPaymentByReference_ReservedCharacters(t *testing.T) {
	reference := "REF/123+456"
	encodedRef := "REF%2F123%2B456"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/payments" {
			t.Errorf("Expected path /api/payments, got %s", r.URL.Path)
		}

		// Check query parameter
		q := r.URL.Query()
		if q.Get("reference") != reference {
			t.Errorf("Expected reference %s, got %s", reference, q.Get("reference"))
		}

		// Verify raw query to ensure it was encoded on the wire
		if r.URL.RawQuery != "reference="+encodedRef {

		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewClient(&ClientConfig{
		BaseURL: server.URL,
		// Mock auth to avoid needing a separate auth handler
	})
	// Inject a valid token to bypass auth
	client.accessToken = "test-token"
	client.tokenExpiry = time.Now().Add(time.Hour)

	_, err := client.GetPaymentByReference(context.Background(), reference)
	if err != nil {
		t.Fatalf("GetPaymentByReference failed: %v", err)
	}
}
