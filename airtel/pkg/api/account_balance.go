package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountBalanceResponse struct {
	Status struct {
		Success     bool   `json:"success"`
		ResultCode  string `json:"result_code"`
		Message     string `json:"message"`
		Code        string `json:"code"`
	} `json:"status"`
	Data struct {
		Balance float64 `json:"balance"`
		Currency string `json:"currency"`
	} `json:"data"`
}

func (s *Service) GetAccountBalance() (*AccountBalanceResponse, error) {
	respBody, err := s.makeRequest(http.MethodGet, accountBalanceURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account balance: %w", err)
	}

	var response AccountBalanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse account balance response: %w", err)
	}

	return &response, nil
}
