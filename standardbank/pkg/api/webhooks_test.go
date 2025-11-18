package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	secret := "test-secret"
	payload := []byte(`{"eventId":"123"}`)

	// Calculate valid signature
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	validSignature := hex.EncodeToString(mac.Sum(nil))

	tests := []struct {
		name      string
		secret    string
		payload   []byte
		signature string
		want      bool
	}{
		{
			name:      "Valid signature",
			secret:    secret,
			payload:   payload,
			signature: validSignature,
			want:      true,
		},
		{
			name:      "Invalid signature",
			secret:    secret,
			payload:   payload,
			signature: "invalid",
			want:      false,
		},
		{
			name:      "Empty secret",
			secret:    "",
			payload:   payload,
			signature: validSignature,
			want:      false,
		},
		{
			name:      "Empty signature",
			secret:    secret,
			payload:   payload,
			signature: "",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wh := NewWebhookHandler(tt.secret)
			// If secret is empty returns nil
			if tt.secret == "" {
				if wh != nil {
					t.Error("Expected nil handler for empty secret")
				}
				// Manually test the static validation function
				if ValidateWebhookSignature(tt.payload, tt.signature, tt.secret) != tt.want {
					t.Errorf("ValidateWebhookSignature() = %v, want %v", !tt.want, tt.want)
				}
				return
			}

			if got := wh.verifySignature(tt.payload, tt.signature); got != tt.want {
				t.Errorf("verifySignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
