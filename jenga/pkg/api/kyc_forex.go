package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	kycEndpoint = "/customer-api/v3.0/identity/verify"
	forexRatesEndpoint = "/forex-api/v3.0/rates"
)

func (c *Client) VerifyIdentity(req KYCRequest) (*KYCResponse, error) {
	if req.DocumentType == "" || req.DocumentNumber == "" || req.CountryCode == "" {
		return nil, fmt.Errorf("documentType, documentNumber, and countryCode are required")
	}

	signatureData := req.DocumentType + req.DocumentNumber + req.CountryCode

	respBody, err := c.SendRequest(http.MethodPost, kycEndpoint, req, signatureData)
	if err != nil {
		return nil, err
	}

	var response KYCResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing KYC response: %w", err)
	}

	return &response, nil
}

func (c *Client) GetForexRates(req ForexRatesRequest) (*ForexRatesResponse, error) {
	if req.CountryCode == "" || req.CurrencyCode == "" {
		return nil, fmt.Errorf("countryCode and currencyCode are required")
	}

	endpoint := fmt.Sprintf("%s/%s/%s", forexRatesEndpoint, req.CountryCode, req.CurrencyCode)
	
	if req.BaseCurrency != "" {
		endpoint = fmt.Sprintf("%s?baseCurrency=%s", endpoint, req.BaseCurrency)
	}

	signatureData := req.CountryCode + req.CurrencyCode

	respBody, err := c.SendRequest(http.MethodGet, endpoint, nil, signatureData)
	if err != nil {
		return nil, err
	}

	var response ForexRatesResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing forex rates response: %w", err)
	}

	return &response, nil
}
