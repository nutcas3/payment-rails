package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PesalinkRequest struct {
	SourceAccount      string  `json:"sourceAccount"`
	DestinationAccount string  `json:"destinationAccount"`
	DestinationBank    string  `json:"destinationBank"` // Bank code
	Amount             float64 `json:"amount"`
	Currency           string  `json:"currency"`
	Reference          string  `json:"reference"`
	Narration          string  `json:"narration"`
	PhoneNumber        string  `json:"phoneNumber"` // Recipient's phone number
}

type PesalinkResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		SourceAccount   string  `json:"sourceAccount"`
		DestAccount     string  `json:"destinationAccount"`
		DestBank        string  `json:"destinationBank"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
	} `json:"data"`
}

type PesalinkStatusRequest struct {
	TransactionID string `json:"transactionId"`
}

type PesalinkStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionID   string  `json:"transactionId"`
		SourceAccount   string  `json:"sourceAccount"`
		DestAccount     string  `json:"destinationAccount"`
		DestBank        string  `json:"destinationBank"`
		Amount          float64 `json:"amount"`
		Currency        string  `json:"currency"`
		TransactionDate string  `json:"transactionDate"`
		Reference       string  `json:"reference"`
		Status          string  `json:"status"`
		StatusReason    string  `json:"statusReason,omitempty"`
	} `json:"data"`
}

func (s *Service) PesalinkTransfer(sourceAccount, destinationAccount, destinationBank string, amount float64, currency, reference, narration, phoneNumber string) (*PesalinkResponse, error) {
	payload := PesalinkRequest{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
		DestinationBank:    destinationBank,
		Amount:             amount,
		Currency:           currency,
		Reference:          reference,
		Narration:          narration,
		PhoneNumber:        phoneNumber,
	}

	respBody, err := s.makeRequest(http.MethodPost, pesalinkURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate PesaLink transfer: %w", err)
	}

	var response PesalinkResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal PesaLink response: %w", err)
	}

	return &response, nil
}

func (s *Service) CheckPesalinkStatus(transactionID string) (*PesalinkStatusResponse, error) {
	payload := PesalinkStatusRequest{
		TransactionID: transactionID,
	}

	respBody, err := s.makeRequest(http.MethodPost, pesalinkStatusURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to check PesaLink status: %w", err)
	}

	var response PesalinkStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal PesaLink status response: %w", err)
	}

	return &response, nil
}
