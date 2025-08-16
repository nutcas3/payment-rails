package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UtilityProvider struct {
	ProviderID   string `json:"providerId"`
	ProviderName string `json:"providerName"`
	Category     string `json:"category"` // e.g., "ELECTRICITY", "WATER", "TV", etc.
	LogoURL      string `json:"logoUrl,omitempty"`
}

type UtilityProvidersResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Providers []UtilityProvider `json:"providers"`
	} `json:"data"`
}

type UtilityPaymentRequest struct {
	SourceAccount string  `json:"sourceAccount"`
	ProviderID    string  `json:"providerId"`
	AccountNumber string  `json:"accountNumber"` // Customer account number with the utility provider
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Reference     string  `json:"reference"`
	PhoneNumber   string  `json:"phoneNumber,omitempty"` // For notifications
}

type UtilityPaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		SourceAccount   string  `json:"sourceAccount"`
		ProviderID      string  `json:"providerId"`
		ProviderName    string  `json:"providerName"`
		AccountNumber   string  `json:"accountNumber"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
		ReceiptNumber   string  `json:"receiptNumber,omitempty"`
	} `json:"data"`
}

type UtilityStatusRequest struct {
	TransactionID string `json:"transactionId"`
}

type UtilityStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		SourceAccount   string  `json:"sourceAccount"`
		ProviderID      string  `json:"providerId"`
		ProviderName    string  `json:"providerName"`
		AccountNumber   string  `json:"accountNumber"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
		StatusReason    string  `json:"statusReason,omitempty"`
		ReceiptNumber   string  `json:"receiptNumber,omitempty"`
	} `json:"data"`
}

func (s *Service) GetUtilityProviders() (*UtilityProvidersResponse, error) {
	respBody, err := s.makeRequest(http.MethodGet, utilityProvidersURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get utility providers: %w", err)
	}

	var response UtilityProvidersResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal utility providers response: %w", err)
	}

	return &response, nil
}

func (s *Service) PayUtility(sourceAccount, providerID, accountNumber string, amount float64, currency, reference, phoneNumber string) (*UtilityPaymentResponse, error) {
	payload := UtilityPaymentRequest{
		SourceAccount: sourceAccount,
		ProviderID:    providerID,
		AccountNumber: accountNumber,
		Amount:        amount,
		Currency:      currency,
		Reference:     reference,
		PhoneNumber:   phoneNumber,
	}

	respBody, err := s.makeRequest(http.MethodPost, utilityPaymentURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to make utility payment: %w", err)
	}

	var response UtilityPaymentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal utility payment response: %w", err)
	}

	return &response, nil
}

func (s *Service) CheckUtilityPaymentStatus(transactionID string) (*UtilityStatusResponse, error) {
	payload := UtilityStatusRequest{
		TransactionID: transactionID,
	}

	respBody, err := s.makeRequest(http.MethodPost, utilityStatusURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to check utility payment status: %w", err)
	}

	var response UtilityStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal utility payment status response: %w", err)
	}

	return &response, nil
}
