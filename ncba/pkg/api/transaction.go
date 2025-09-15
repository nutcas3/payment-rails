package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TransactionStatus struct {
	TransactionID   string  `json:"transactionId"`
	Status         string  `json:"status"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	Reference      string  `json:"reference"`
	Timestamp      string  `json:"timestamp"`
	Type           string  `json:"type"`
	SourceAccount  string  `json:"sourceAccount"`
	Description    string  `json:"description"`
	ResponseCode   string  `json:"responseCode"`
	ResponseDesc   string  `json:"responseDesc"`
}

func (c *Client) CheckTransactionStatus(transactionID string) (*TransactionStatus, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/transactions/%s/status", BaseURL, transactionID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result TransactionStatus
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}
