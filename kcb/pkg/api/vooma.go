package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VoomaPayRequest struct {
	Amount float64 `json:"amount"`
}

type VoomaPayResponse struct {
	Status struct {
		Success    bool   `json:"success"`
		ResultCode string `json:"result_code"`
		Message    string `json:"message"`
		Code       string `json:"code"`
	} `json:"status"`
	Data struct {
		TransactionID   string  `json:"transaction_id"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		Status          string  `json:"status"`
		TransactionDate string  `json:"transaction_date"`
		Reference       string  `json:"reference"`
	} `json:"data"`
}

func (s *Service) VoomaPay(amount float64) (*VoomaPayResponse, error) {
	payload := VoomaPayRequest{
		Amount: amount,
	}

	respBody, err := s.makeRequest(http.MethodPost, voomaPayURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate Vooma payment: %w", err)
	}

	var response VoomaPayResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse Vooma payment response: %w", err)
	}

	return &response, nil
}
