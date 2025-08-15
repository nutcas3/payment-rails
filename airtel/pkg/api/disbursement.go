package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DisbursementRequest struct {
	Reference  string `json:"reference"`
	PIN        string `json:"pin,omitempty"`
	Subscriber struct {
		Country  string `json:"country"`
		Currency string `json:"currency"`
		MSISDN   string `json:"msisdn"`
	} `json:"subscriber"`
	Transaction struct {
		Amount     float64 `json:"amount"`
		ID         string  `json:"id"`
		Reference  string  `json:"reference"`
	} `json:"transaction"`
}

type DisbursementResponse struct {
	Status struct {
		Success     bool   `json:"success"`
		ResultCode  string `json:"result_code"`
		Message     string `json:"message"`
		Code        string `json:"code"`
	} `json:"status"`
	Data struct {
		Transaction struct {
			ID        string `json:"id"`
			Status    string `json:"status"`
			AirtelMoney struct {
				ID        string `json:"id"`
				Status    string `json:"status"`
				Message   string `json:"message"`
			} `json:"airtel_money"`
		} `json:"transaction"`
	} `json:"data"`
}

type DisbursementStatusResponse struct {
	Status struct {
		Success     bool   `json:"success"`
		ResultCode  string `json:"result_code"`
		Message     string `json:"message"`
		Code        string `json:"code"`
	} `json:"status"`
	Data struct {
		Transaction struct {
			ID        string `json:"id"`
			Status    string `json:"status"`
			AirtelMoney struct {
				ID        string `json:"id"`
				Status    string `json:"status"`
				Message   string `json:"message"`
			} `json:"airtel_money"`
		} `json:"transaction"`
	} `json:"data"`
}

func (s *Service) Disburse(reference, phone string, amount float64, transactionID string, pin string) (*DisbursementResponse, error) {
	if reference == "" {
		return nil, fmt.Errorf("reference is required")
	}
	if phone == "" {
		return nil, fmt.Errorf("phone number is required")
	}
	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}
	if transactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	if len(phone) > 9 && phone[0:1] == "+" {
		phone = phone[1:] // Remove the + sign
	}
	
	req := DisbursementRequest{
		Reference: reference,
		PIN:       pin,
	}
	req.Subscriber.Country = s.country
	req.Subscriber.Currency = s.currency
	req.Subscriber.MSISDN = phone
	req.Transaction.Amount = amount
	req.Transaction.ID = transactionID
	req.Transaction.Reference = reference

	respBody, err := s.makeRequest(http.MethodPost, disburseURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate disbursement: %w", err)
	}

	var response DisbursementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse disbursement response: %w", err)
	}

	return &response, nil
}

func (s *Service) GetDisbursementStatus(transactionID string) (*DisbursementStatusResponse, error) {
	if transactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	url := disbursementStatusURL + transactionID
	respBody, err := s.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get disbursement status: %w", err)
	}

	var response DisbursementStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse disbursement status response: %w", err)
	}

	return &response, nil
}
