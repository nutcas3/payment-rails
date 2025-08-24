package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	billPaymentEndpoint = "/transaction-api/v3.0/bills/pay"
)

func (c *Client) PayBill(req BillPaymentRequest) (*BillPaymentResponse, error) {
	if req.BillerCode == "" || req.AccountNumber == "" || req.Amount == "" || 
	   req.Reference == "" || req.CurrencyCode == "" {
		return nil, fmt.Errorf("missing required fields in BillPaymentRequest")
	}

	signatureData := req.BillerCode + req.AccountNumber + req.Amount + req.Reference

	respBody, err := c.SendRequest(http.MethodPost, billPaymentEndpoint, req, signatureData)
	if err != nil {
		return nil, err
	}

	var response BillPaymentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing bill payment response: %w", err)
	}

	return &response, nil
}
