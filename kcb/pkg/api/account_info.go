package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountInfoResponse struct {
	Status struct {
		Success    bool   `json:"success"`
		ResultCode string `json:"result_code"`
		Message    string `json:"message"`
		Code       string `json:"code"`
	} `json:"status"`
	Data struct {
		AccountNumber string  `json:"account_number"`
		AccountName   string  `json:"account_name"`
		Balance       float64 `json:"balance"`
		Currency      string  `json:"currency"`
		AccountType   string  `json:"account_type"`
		Branch        string  `json:"branch"`
		Status        string  `json:"status"`
	} `json:"data"`
}

func (s *Service) GetAccountInfo() (*AccountInfoResponse, error) {
	respBody, err := s.makeRequest(http.MethodGet, accountInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account information: %w", err)
	}

	var response AccountInfoResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse account information response: %w", err)
	}

	return &response, nil
}
