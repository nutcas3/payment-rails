package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TransferRequest struct {
	SourceAccount      string  `json:"sourceAccount"`
	DestinationAccount string  `json:"destinationAccount"`
	Amount             float64 `json:"amount"`
	Currency           string  `json:"currency"`
	Reference          string  `json:"reference"`
	Narration          string  `json:"narration"`
}

type TransferResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		SourceAccount   string  `json:"sourceAccount"`
		DestAccount     string  `json:"destinationAccount"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
	} `json:"data"`
}

func (s *Service) TransferFunds(sourceAccount, destinationAccount string, amount float64, currency, reference, narration string) (*TransferResponse, error) {
	payload := TransferRequest{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
		Amount:             amount,
		Currency:           currency,
		Reference:          reference,
		Narration:          narration,
	}

	respBody, err := s.makeRequest(http.MethodPost, accountTransferURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer funds: %w", err)
	}

	var response TransferResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transfer response: %w", err)
	}

	return &response, nil
}
