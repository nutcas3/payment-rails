package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ForexRatesResponse struct {
	Status struct {
		Success    bool   `json:"success"`
		ResultCode string `json:"result_code"`
		Message    string `json:"message"`
		Code       string `json:"code"`
	} `json:"status"`
	Data struct {
		BaseCurrency string `json:"base_currency"`
		Rates        map[string]float64 `json:"rates"`
		Timestamp    string `json:"timestamp"`
	} `json:"data"`
}

type ForexExchangeResponse struct {
	Status struct {
		Success    bool   `json:"success"`
		ResultCode string `json:"result_code"`
		Message    string `json:"message"`
		Code       string `json:"code"`
	} `json:"status"`
	Data struct {
		FromCurrency string  `json:"from_currency"`
		ToCurrency   string  `json:"to_currency"`
		Amount       float64 `json:"amount"`
		ConvertedAmount float64 `json:"converted_amount"`
		ExchangeRate float64 `json:"exchange_rate"`
		Timestamp    string  `json:"timestamp"`
	} `json:"data"`
}

type ForexExchangeRequest struct {
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
	Amount       float64 `json:"amount"`
}

func (s *Service) GetForexRates(currency string) (*ForexRatesResponse, error) {
	url := fmt.Sprintf("%s/%s", forexRatesURL, currency)
	respBody, err := s.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get forex rates: %w", err)
	}

	var response ForexRatesResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse forex rates response: %w", err)
	}

	return &response, nil
}

func (s *Service) ExchangeCurrency(from, to string, amount float64) (*ForexExchangeResponse, error) {
	payload := ForexExchangeRequest{
		FromCurrency: from,
		ToCurrency:   to,
		Amount:       amount,
	}

	respBody, err := s.makeRequest(http.MethodPost, forexExchangeURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange currency: %w", err)
	}

	var response ForexExchangeResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse currency exchange response: %w", err)
	}

	return &response, nil
}
