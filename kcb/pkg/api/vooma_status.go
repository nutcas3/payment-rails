package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VoomaStatusRequest struct {
	TransactionID string `json:"transactionId"`
}

type VoomaStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
		StatusReason    string  `json:"statusReason,omitempty"`
	} `json:"data"`
}

func (s *Service) CheckVoomaStatus(transactionID string) (*VoomaStatusResponse, error) {
	payload := VoomaStatusRequest{
		TransactionID: transactionID,
	}

	respBody, err := s.makeRequest(http.MethodPost, voomaStatusURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to check Vooma payment status: %w", err)
	}

	var response VoomaStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Vooma status response: %w", err)
	}

	return &response, nil
}
