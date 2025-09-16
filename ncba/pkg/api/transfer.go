package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TransferRequest struct {
	SourceAccount      string  `json:"sourceAccount"`
	DestinationAccount string  `json:"destinationAccount"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	Reference         string  `json:"reference"`
	Narration         string  `json:"narration"`
}

type InternalTransferRequest struct {
	TransferRequest
	DestinationName string `json:"destinationName"`
}

type ExternalTransferRequest struct {
	TransferRequest
	BankCode         string `json:"bankCode"`
	BranchCode       string `json:"branchCode"`
	DestinationName  string `json:"destinationName"`
}

type RTGSTransferRequest struct {
	ExternalTransferRequest
	SwiftCode string `json:"swiftCode"`
}

type PesaLinkTransferRequest struct {
	TransferRequest
	DestinationBank string `json:"destinationBank"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
}

type TransferResponse struct {
	TransactionID string `json:"transactionId"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}

func (c *Client) SendInternalTransfer(req InternalTransferRequest) (*TransferResponse, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest("POST", BaseURL+"/transfers/internal", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(httpReq)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result TransferResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}

func (c *Client) SendExternalTransfer(req ExternalTransferRequest) (*TransferResponse, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest("POST", BaseURL+"/transfers/external", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(httpReq)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result TransferResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}

func (c *Client) SendRTGSTransfer(req RTGSTransferRequest) (*TransferResponse, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest("POST", BaseURL+"/transfers/rtgs", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(httpReq)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result TransferResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}

func (c *Client) SendPesaLinkTransfer(req PesaLinkTransferRequest) (*TransferResponse, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest("POST", BaseURL+"/transfers/pesalink", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(httpReq)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result TransferResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}
