package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountBalanceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AccountNumber string  `json:"accountNumber"`
		AccountName   string  `json:"accountName"`
		Balance       float64 `json:"balance"`
		Currency      string  `json:"currency"`
		AsOf          string  `json:"asOf"`
	} `json:"data"`
}

func (s *Service) GetAccountBalance(accountNumber string) (*AccountBalanceResponse, error) {
	url := fmt.Sprintf("%s?accountNumber=%s", accountBalanceURL, accountNumber)
	
	respBody, err := s.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account balance: %w", err)
	}

	var response AccountBalanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account balance response: %w", err)
	}

	return &response, nil
}
