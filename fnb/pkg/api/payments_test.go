package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupTestServer(t *testing.T, handler http.HandlerFunc) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth/token" {
			resp := AuthResponse{
				AccessToken: "mock-token",
				TokenType:   "Bearer",
				ExpiresIn:   3600,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		handler(w, r)
	}))

	config := &ClientConfig{
		ClientID:     "test-client",
		ClientSecret: "test-secret",
		APIKey:       "test-key",
		BaseURL:      server.URL,
	}

	client := NewClient(config)
	return server, client
}

func TestCreateEFTPayment(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/payments/eft" {
			t.Errorf("Expected path /api/v1/payments/eft, got %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		var req EFTPaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req.Amount != 1000.00 {
			t.Errorf("Expected amount 1000.00, got %.2f", req.Amount)
		}

		resp := EFTPaymentResponse{
			TransactionID:     "TXN123456",
			Status:            "PENDING",
			StatusDescription: "Payment submitted successfully",
			PaymentReference:  req.PaymentReference,
			Amount:            req.Amount,
			Currency:          req.Currency,
			ProcessingDate:    time.Now(),
			BeneficiaryName:   req.BeneficiaryName,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	ctx := context.Background()
	req := EFTPaymentRequest{
		SourceAccountNumber:      "1234567890",
		BeneficiaryAccountNumber: "0987654321",
		BeneficiaryName:          "John Doe",
		BeneficiaryBankCode:      "250655",
		Amount:                   1000.00,
		Currency:                 "ZAR",
		PaymentReference:         "PAY001",
		PaymentDescription:       "Test payment",
	}

	resp, err := client.CreateEFTPayment(ctx, req)
	if err != nil {
		t.Fatalf("CreateEFTPayment failed: %v", err)
	}

	if resp.TransactionID != "TXN123456" {
		t.Errorf("Expected transaction ID 'TXN123456', got %s", resp.TransactionID)
	}

	if resp.Status != "PENDING" {
		t.Errorf("Expected status 'PENDING', got %s", resp.Status)
	}

	if resp.Amount != 1000.00 {
		t.Errorf("Expected amount 1000.00, got %.2f", resp.Amount)
	}
}

func TestGetPaymentStatus(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/payments/TXN123456/status" {
			t.Errorf("Expected path /api/v1/payments/TXN123456/status, got %s", r.URL.Path)
		}

		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		resp := PaymentStatusResponse{
			TransactionID:      "TXN123456",
			PaymentReference:   "PAY001",
			Status:             "COMPLETED",
			StatusDescription:  "Payment completed successfully",
			Amount:             1000.00,
			Currency:           "ZAR",
			SourceAccount:      "1234567890",
			BeneficiaryAccount: "0987654321",
			BeneficiaryName:    "John Doe",
			ProcessingDate:     time.Now(),
			SettlementDate:     time.Now().Format("2006-01-02"),
			LastUpdated:        time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	ctx := context.Background()
	resp, err := client.GetPaymentStatus(ctx, "TXN123456")
	if err != nil {
		t.Fatalf("GetPaymentStatus failed: %v", err)
	}

	if resp.TransactionID != "TXN123456" {
		t.Errorf("Expected transaction ID 'TXN123456', got %s", resp.TransactionID)
	}

	if resp.Status != "COMPLETED" {
		t.Errorf("Expected status 'COMPLETED', got %s", resp.Status)
	}
}

func TestCreateBatchPayment(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/payments/batch" {
			t.Errorf("Expected path /api/v1/payments/batch, got %s", r.URL.Path)
		}

		var req BatchPaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if len(req.Payments) != 2 {
			t.Errorf("Expected 2 payments, got %d", len(req.Payments))
		}

		resp := BatchPaymentResponse{
			BatchID:        "BATCH001",
			BatchReference: req.BatchReference,
			Status:         "PROCESSING",
			TotalAmount:    req.TotalAmount,
			TotalCount:     req.TotalCount,
			SuccessCount:   0,
			FailureCount:   0,
			ProcessingDate: time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	ctx := context.Background()
	req := BatchPaymentRequest{
		BatchReference:      "BATCH001",
		SourceAccountNumber: "1234567890",
		TotalAmount:         3000.00,
		TotalCount:          2,
		Payments: []EFTPaymentRequest{
			{
				BeneficiaryAccountNumber: "1111111111",
				BeneficiaryName:          "Beneficiary 1",
				BeneficiaryBankCode:      "250655",
				Amount:                   1500.00,
				PaymentReference:         "PAY001",
			},
			{
				BeneficiaryAccountNumber: "2222222222",
				BeneficiaryName:          "Beneficiary 2",
				BeneficiaryBankCode:      "250655",
				Amount:                   1500.00,
				PaymentReference:         "PAY002",
			},
		},
	}

	resp, err := client.CreateBatchPayment(ctx, req)
	if err != nil {
		t.Fatalf("CreateBatchPayment failed: %v", err)
	}

	if resp.BatchID != "BATCH001" {
		t.Errorf("Expected batch ID 'BATCH001', got %s", resp.BatchID)
	}

	if resp.TotalCount != 2 {
		t.Errorf("Expected total count 2, got %d", resp.TotalCount)
	}
}

func TestCancelPayment(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/payments/TXN123456/cancel" {
			t.Errorf("Expected path /api/v1/payments/TXN123456/cancel, got %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	ctx := context.Background()
	err := client.CancelPayment(ctx, "TXN123456", "Customer request")
	if err != nil {
		t.Fatalf("CancelPayment failed: %v", err)
	}
}
