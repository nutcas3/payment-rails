package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type StatementRequest struct {
	AccountNumber string `json:"accountNumber"`
	StartDate     string `json:"startDate"` // Format: YYYY-MM-DD
	EndDate       string `json:"endDate"`   // Format: YYYY-MM-DD
}

type Transaction struct {
	TransactionID   string    `json:"transactionId"`
	TransactionDate time.Time `json:"transactionDate"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Type            string    `json:"type"` // "DEBIT" or "CREDIT"
	Balance         float64   `json:"balance"`
	Reference       string    `json:"reference"`
}

type StatementResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AccountNumber string        `json:"accountNumber"`
		AccountName   string        `json:"accountName"`
		StartDate     string        `json:"startDate"`
		EndDate       string        `json:"endDate"`
		Currency      string        `json:"currency"`
		Transactions  []Transaction `json:"transactions"`
	} `json:"data"`
}

func (s *Service) GetAccountStatement(accountNumber, startDate, endDate string) (*StatementResponse, error) {
	payload := StatementRequest{
		AccountNumber: accountNumber,
		StartDate:     startDate,
		EndDate:       endDate,
	}

	respBody, err := s.makeRequest(http.MethodPost, accountStatementURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to get account statement: %w", err)
	}

	var response StatementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account statement response: %w", err)
	}

	return &response, nil
}
