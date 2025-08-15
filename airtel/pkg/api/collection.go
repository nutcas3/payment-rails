package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CollectionRequest struct {
	Reference  string `json:"reference"`
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

type CollectionResponse struct {
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

type TransactionStatusResponse struct {
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

type RefundRequest struct {
	Transaction struct {
		AirtelMoneyID string  `json:"airtel_money_id"`
		Amount        float64 `json:"amount"`
	} `json:"transaction"`
}

type RefundResponse struct {
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
		} `json:"transaction"`
	} `json:"data"`
}

func (s *Service) UssdPush(reference, phone string, amount float64, transactionID string) (*CollectionResponse, error) {
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
	
	req := CollectionRequest{
		Reference: reference,
	}
	req.Subscriber.Country = s.country
	req.Subscriber.Currency = s.currency
	req.Subscriber.MSISDN = phone
	req.Transaction.Amount = amount
	req.Transaction.ID = transactionID
	req.Transaction.Reference = reference

	respBody, err := s.makeRequest(http.MethodPost, ussdPushURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate USSD Push: %w", err)
	}

	var response CollectionResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse USSD Push response: %w", err)
	}

	return &response, nil
}

func (s *Service) GetTransactionStatus(transactionID string) (*TransactionStatusResponse, error) {
	if transactionID == "" {
		return nil, fmt.Errorf("transaction ID is required")
	}

	url := transactionStatusURL + transactionID
	respBody, err := s.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction status: %w", err)
	}

	var response TransactionStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse transaction status response: %w", err)
	}

	return &response, nil
}

func (s *Service) RefundTransaction(airtelMoneyID string, amount float64) (*RefundResponse, error) {
	if airtelMoneyID == "" {
		return nil, fmt.Errorf("airtel Money ID is required")
	}
	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	req := RefundRequest{}
	req.Transaction.AirtelMoneyID = airtelMoneyID
	req.Transaction.Amount = amount

	respBody, err := s.makeRequest(http.MethodPost, refundURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate refund: %w", err)
	}

	var response RefundResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse refund response: %w", err)
	}

	return &response, nil
}
