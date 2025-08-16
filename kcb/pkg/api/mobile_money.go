package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MobileMoneyRequest struct {
	SourceAccount string  `json:"sourceAccount"`
	PhoneNumber   string  `json:"phoneNumber"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Reference     string  `json:"reference"`
	Narration     string  `json:"narration"`
	Provider      string  `json:"provider"` // e.g., "MPESA", "AIRTEL", etc.
}

type MobileMoneyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		SourceAccount   string  `json:"sourceAccount"`
		PhoneNumber     string  `json:"phoneNumber"`
		Provider        string  `json:"provider"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
	} `json:"data"`
}

type MobileMoneyStatusRequest struct {
	TransactionID string `json:"transactionId"`
}

type MobileMoneyStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		SourceAccount   string  `json:"sourceAccount"`
		PhoneNumber     string  `json:"phoneNumber"`
		Provider        string  `json:"provider"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
		StatusReason    string  `json:"statusReason,omitempty"`
	} `json:"data"`
}

func (s *Service) MobileMoneyTransfer(sourceAccount, phoneNumber string, amount float64, currency, reference, narration, provider string) (*MobileMoneyResponse, error) {
	payload := MobileMoneyRequest{
		SourceAccount: sourceAccount,
		PhoneNumber:   phoneNumber,
		Amount:        amount,
		Currency:      currency,
		Reference:     reference,
		Narration:     narration,
		Provider:      provider,
	}

	respBody, err := s.makeRequest(http.MethodPost, mobileMoneyURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate mobile money transfer: %w", err)
	}

	var response MobileMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mobile money response: %w", err)
	}

	return &response, nil
}

func (s *Service) CheckMobileMoneyStatus(transactionID string) (*MobileMoneyStatusResponse, error) {
	payload := MobileMoneyStatusRequest{
		TransactionID: transactionID,
	}

	respBody, err := s.makeRequest(http.MethodPost, mobileMoneyStatusURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to check mobile money status: %w", err)
	}

	var response MobileMoneyStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mobile money status response: %w", err)
	}

	return &response, nil
}
