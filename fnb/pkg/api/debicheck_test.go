package api

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestCreateMandate(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/debicheck/mandates" {
			t.Errorf("Expected path /api/v1/debicheck/mandates, got %s", r.URL.Path)
		}

		var req MandateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req.MaximumAmount != 5000.00 {
			t.Errorf("Expected maximum amount 5000.00, got %.2f", req.MaximumAmount)
		}

		resp := MandateResponse{
			MandateID:           "MAN123456",
			Status:              "PENDING_APPROVAL",
			StatusDescription:   "Mandate awaiting customer approval",
			ContractReference:   req.ContractReference,
			DebtorName:          req.DebtorName,
			MaximumAmount:       req.MaximumAmount,
			Currency:            req.Currency,
			FrequencyType:       req.FrequencyType,
			FirstCollectionDate: req.FirstCollectionDate,
			CreatedDate:         time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	ctx := context.Background()
	req := MandateRequest{
		CreditorName:          "Test Company",
		CreditorAbbreviation:  "TESTCO",
		CreditorAccountNumber: "1234567890",
		DebtorName:            "John Doe",
		DebtorIDNumber:        "8001015009087",
		DebtorAccountNumber:   "9876543210",
		DebtorBankCode:        "250655",
		ContractReference:     "CONTRACT001",
		MaximumAmount:         5000.00,
		Currency:              "ZAR",
		FrequencyType:         "MONTHLY",
		FirstCollectionDate:   "2025-11-01",
		MandateDescription:    "Monthly subscription payment",
	}

	resp, err := client.CreateMandate(ctx, req)
	if err != nil {
		t.Fatalf("CreateMandate failed: %v", err)
	}

	if resp.MandateID != "MAN123456" {
		t.Errorf("Expected mandate ID 'MAN123456', got %s", resp.MandateID)
	}

	if resp.Status != "PENDING_APPROVAL" {
		t.Errorf("Expected status 'PENDING_APPROVAL', got %s", resp.Status)
	}

	if resp.MaximumAmount != 5000.00 {
		t.Errorf("Expected maximum amount 5000.00, got %.2f", resp.MaximumAmount)
	}
}

func TestGetMandateStatus(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/debicheck/mandates/MAN123456" {
			t.Errorf("Expected path /api/v1/debicheck/mandates/MAN123456, got %s", r.URL.Path)
		}

		resp := MandateStatusResponse{
			MandateID:           "MAN123456",
			ContractReference:   "CONTRACT001",
			Status:              "ACTIVE",
			StatusDescription:   "Mandate is active",
			CreditorName:        "Test Company",
			DebtorName:          "John Doe",
			DebtorAccountNumber: "9876543210",
			MaximumAmount:       5000.00,
			Currency:            "ZAR",
			FrequencyType:       "MONTHLY",
			FirstCollectionDate: "2025-11-01",
			NextCollectionDate:  "2025-11-01",
			CreatedDate:         time.Now().Add(-24 * time.Hour),
			ApprovalDate:        time.Now().Format("2006-01-02"),
			LastModifiedDate:    time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	ctx := context.Background()
	resp, err := client.GetMandateStatus(ctx, "MAN123456")
	if err != nil {
		t.Fatalf("GetMandateStatus failed: %v", err)
	}

	if resp.MandateID != "MAN123456" {
		t.Errorf("Expected mandate ID 'MAN123456', got %s", resp.MandateID)
	}

	if resp.Status != "ACTIVE" {
		t.Errorf("Expected status 'ACTIVE', got %s", resp.Status)
	}
}

func TestCollectAgainstMandate(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/debicheck/collect" {
			t.Errorf("Expected path /api/v1/debicheck/collect, got %s", r.URL.Path)
		}

		var req MandateCollectionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req.Amount != 1000.00 {
			t.Errorf("Expected amount 1000.00, got %.2f", req.Amount)
		}

		resp := EFTCollectionResponse{
			TransactionID:       "TXN789012",
			Status:              "PENDING",
			StatusDescription:   "Collection submitted",
			CollectionReference: req.CollectionReference,
			Amount:              req.Amount,
			Currency:            "ZAR",
			ProcessingDate:      time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	ctx := context.Background()
	req := MandateCollectionRequest{
		MandateID:           "MAN123456",
		Amount:              1000.00,
		CollectionReference: "COLL001",
		Description:         "Monthly subscription",
	}

	resp, err := client.CollectAgainstMandate(ctx, req)
	if err != nil {
		t.Fatalf("CollectAgainstMandate failed: %v", err)
	}

	if resp.TransactionID != "TXN789012" {
		t.Errorf("Expected transaction ID 'TXN789012', got %s", resp.TransactionID)
	}

	if resp.Status != "PENDING" {
		t.Errorf("Expected status 'PENDING', got %s", resp.Status)
	}
}

func TestCancelMandate(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/debicheck/mandates/MAN123456/cancel" {
			t.Errorf("Expected path /api/v1/debicheck/mandates/MAN123456/cancel, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	ctx := context.Background()
	req := MandateCancellationRequest{
		MandateID:          "MAN123456",
		CancellationReason: "Customer request",
	}

	err := client.CancelMandate(ctx, req)
	if err != nil {
		t.Fatalf("CancelMandate failed: %v", err)
	}
}

func TestVerifyMandate(t *testing.T) {
	server, client := setupTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/debicheck/mandates/verify" {
			t.Errorf("Expected path /api/v1/debicheck/mandates/verify, got %s", r.URL.Path)
		}

		resp := map[string]interface{}{
			"valid":   true,
			"message": "Mandate is valid for collection",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	ctx := context.Background()
	valid, err := client.VerifyMandate(ctx, "MAN123456", 1000.00)
	if err != nil {
		t.Fatalf("VerifyMandate failed: %v", err)
	}

	if !valid {
		t.Error("Expected mandate to be valid")
	}
}
