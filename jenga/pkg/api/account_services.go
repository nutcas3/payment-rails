package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	accountBalanceEndpoint = "/account-api/v3.0/accounts/balances"
	miniStatementEndpoint  = "/account-api/v3.0/accounts/miniStatement"
	fullStatementEndpoint  = "/account-api/v3.0/accounts/fullStatement"
)

func (c *Client) GetAccountBalance(req AccountBalanceRequest) (*AccountBalanceResponse, error) {
	if req.CountryCode == "" || req.AccountID == "" {
		return nil, fmt.Errorf("countryCode and accountId are required")
	}

	endpoint := fmt.Sprintf("%s/%s/%s", accountBalanceEndpoint, req.CountryCode, req.AccountID)
	
	queryParams := make([]string, 0)
	if req.AccountType != "" {
		queryParams = append(queryParams, fmt.Sprintf("accountType=%s", req.AccountType))
	}
	if req.CurrencyCode != "" {
		queryParams = append(queryParams, fmt.Sprintf("currencyCode=%s", req.CurrencyCode))
	}
	
	if len(queryParams) > 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, joinQueryParams(queryParams))
	}

	signatureData := req.AccountID

	respBody, err := c.SendRequest(http.MethodGet, endpoint, nil, signatureData)
	if err != nil {
		return nil, err
	}

	var response AccountBalanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing account balance response: %w", err)
	}

	return &response, nil
}

func (c *Client) GetMiniStatement(req MiniStatementRequest) (*MiniStatementResponse, error) {
	if req.CountryCode == "" || req.AccountID == "" {
		return nil, fmt.Errorf("countryCode and accountId are required")
	}


	endpoint := fmt.Sprintf("%s/%s/%s", miniStatementEndpoint, req.CountryCode, req.AccountID)
	
	queryParams := make([]string, 0)
	if req.AccountType != "" {
		queryParams = append(queryParams, fmt.Sprintf("accountType=%s", req.AccountType))
	}
	if req.CurrencyCode != "" {
		queryParams = append(queryParams, fmt.Sprintf("currencyCode=%s", req.CurrencyCode))
	}
	
	if len(queryParams) > 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, joinQueryParams(queryParams))
	}

	// Updated signature generation according to API documentation
	signatureData := req.CountryCode + req.AccountID

	respBody, err := c.SendRequest(http.MethodGet, endpoint, nil, signatureData)
	if err != nil {
		return nil, err
	}

	var response MiniStatementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing mini statement response: %w", err)
	}

	return &response, nil
}

func (c *Client) GetFullStatement(req FullStatementRequest) (*FullStatementResponse, error) {
	if req.CountryCode == "" || req.AccountID == "" || req.FromDate == "" || req.ToDate == "" {
		return nil, fmt.Errorf("countryCode, accountId, fromDate, and toDate are required")
	}

	// Updated to use POST method with request body instead of query parameters
	endpoint := fullStatementEndpoint
	
	// Create request body
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	// Updated signature generation according to API documentation
	signatureData := req.AccountID + req.CountryCode + req.ToDate

	respBody, err := c.SendRequest(http.MethodPost, endpoint, requestBody, signatureData)
	if err != nil {
		return nil, err
	}

	var response FullStatementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing full statement response: %w", err)
	}

	return &response, nil
}

func joinQueryParams(params []string) string {
	return strings.Join(params, "&")
}
