package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	kycEndpoint = "/customer-api/v3.0/identity/verify"
	amlScreeningEndpoint = "/customer-api/v3.0/aml/screening"
	cddEndpoint = "/customer-api/v3.0/cdd/verify"
	forexRatesEndpoint = "/forex-api/v3.0/rates"
)

func (c *Client) VerifyIdentity(req KYCRequest) (*KYCResponse, error) {
	if req.DocumentType == "" || req.DocumentNumber == "" || req.CountryCode == "" {
		return nil, fmt.Errorf("documentType, documentNumber, and countryCode are required")
	}

	// Create request body
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	signatureData := req.DocumentType + req.DocumentNumber + req.CountryCode

	respBody, err := c.SendRequest(http.MethodPost, kycEndpoint, requestBody, signatureData)
	if err != nil {
		return nil, err
	}

	var response KYCResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing KYC response: %w", err)
	}

	return &response, nil
}

// PerformAMLScreening performs Anti-Money Laundering screening on an individual
func (c *Client) PerformAMLScreening(req AMLScreeningRequest) (*AMLScreeningResponse, error) {
	if req.FirstName == "" || req.LastName == "" || req.CountryCode == "" {
		return nil, fmt.Errorf("firstName, lastName, and countryCode are required")
	}

	// Create request body
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	// Generate signature according to API documentation
	signatureData := req.FirstName + req.LastName + req.CountryCode

	// Send POST request
	respBody, err := c.SendRequest(http.MethodPost, amlScreeningEndpoint, requestBody, signatureData)
	if err != nil {
		return nil, err
	}

	var response AMLScreeningResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing AML screening response: %w", err)
	}

	return &response, nil
}

// PerformCustomerDueDiligence performs Customer Due Diligence (CDD) checks
func (c *Client) PerformCustomerDueDiligence(req CDDRequest) (*CDDResponse, error) {
	if req.CustomerID == "" || req.CountryCode == "" {
		return nil, fmt.Errorf("customerID and countryCode are required")
	}

	// Create request body
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	// Generate signature according to API documentation
	signatureData := req.CustomerID + req.CountryCode

	// Send POST request
	respBody, err := c.SendRequest(http.MethodPost, cddEndpoint, requestBody, signatureData)
	if err != nil {
		return nil, err
	}

	var response CDDResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing CDD response: %w", err)
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
