package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	billPaymentEndpoint = "/transaction-api/v3.0/bills/pay"
	airtimeEndpoint = "/transaction-api/v3.0/airtime"
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

func (c *Client) PurchaseAirtime(req AirtimePurchaseRequest) (*AirtimePurchaseResponse, error) {
	if req.CustomerMobile == "" || req.TelcoCode == "" || req.Amount == "" || 
	   req.Reference == "" || req.CurrencyCode == "" {
		return nil, fmt.Errorf("missing required fields in AirtimePurchaseRequest")
	}

	signatureData := req.CustomerMobile + req.TelcoCode + req.Amount + req.Reference

	respBody, err := c.SendRequest(http.MethodPost, airtimeEndpoint, req, signatureData)
	if err != nil {
		return nil, err
	}

	var response AirtimePurchaseResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing airtime purchase response: %w", err)
	}

	return &response, nil
}
