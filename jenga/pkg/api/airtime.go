package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	airtimePurchaseEndpoint = "/transaction-api/v3.0/airtime"
)

func (c *Client) PurchaseAirtime(req AirtimePurchaseRequest) (*AirtimePurchaseResponse, error) {
	if req.CustomerMobile == "" || req.TelcoCode == "" || req.Amount == "" || req.Reference == "" || req.CurrencyCode == "" {
		return nil, fmt.Errorf("customerMobile, telcoCode, amount, reference, and currencyCode are required")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	signatureData := req.CustomerMobile + req.TelcoCode + req.Amount + req.CurrencyCode

	respBody, err := c.SendRequest(http.MethodPost, airtimePurchaseEndpoint, requestBody, signatureData)
	if err != nil {
		return nil, err
	}

	var response AirtimePurchaseResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing airtime purchase response: %w", err)
	}

	return &response, nil
}
