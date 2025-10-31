package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	remittanceTransferURL     = "/remittance/v1_0/transfer"
	remittanceStatusURL       = "/remittance/v1_0/transfer/%s"
	remittanceBalanceURL      = "/remittance/v1_0/account/balance"
	remittanceActiveURL       = "/remittance/v1_0/accountholder/msisdn/%s/active"
	remittanceUserInfoURL     = "/remittance/v1_0/accountholder/msisdn/%s/basicuserinfo"
)

type RemittanceRequest struct {
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	ExternalID        string `json:"externalId"`
	Payee             Payee  `json:"payee"`
	PayerMessage      string `json:"payerMessage"`
	PayeeNote         string `json:"payeeNote"`
	CallbackURL       string `json:"callbackUrl,omitempty"`
}

// RemittanceResponse represents the response from remittance
type RemittanceResponse struct {
	ReferenceID string `json:"referenceId"`
	Status      string `json:"status"`
	Reason      string `json:"reason,omitempty"`
}

// Remit initiates a remittance transfer
func (c *Client) Remit(req RemittanceRequest) (*RemittanceResponse, error) {
	// Generate reference ID
	referenceID := generateUUID()

	headers := map[string]string{
		"X-Reference-Id": referenceID,
	}

	_, err := c.makeRequest(http.MethodPost, remittanceTransferURL, "remittance", req, headers)
	if err != nil {
		return nil, fmt.Errorf("remittance failed: %w", err)
	}

	// For successful request, return the reference ID
	return &RemittanceResponse{
		ReferenceID: referenceID,
		Status:      "PENDING",
	}, nil
}

// GetRemittanceStatus retrieves the status of a remittance
func (c *Client) GetRemittanceStatus(referenceID string) (*TransactionStatus, error) {
	url := fmt.Sprintf(remittanceStatusURL, referenceID)

	respBody, err := c.makeRequest(http.MethodGet, url, "remittance", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get remittance status: %w", err)
	}

	var status TransactionStatus
	if err := json.Unmarshal(respBody, &status); err != nil {
		return nil, fmt.Errorf("failed to decode remittance status: %w", err)
	}

	return &status, nil
}

func (c *Client) GetRemittanceBalance() (*Balance, error) {
	respBody, err := c.makeRequest(http.MethodGet, remittanceBalanceURL, "remittance", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get remittance balance: %w", err)
	}

	var balance Balance
	if err := json.Unmarshal(respBody, &balance); err != nil {
		return nil, fmt.Errorf("failed to decode balance: %w", err)
	}

	return &balance, nil
}

func (c *Client) ValidateRemittanceAccountHolder(msisdn string) (*AccountHolder, error) {
	url := fmt.Sprintf(remittanceActiveURL, msisdn)

	respBody, err := c.makeRequest(http.MethodGet, url, "remittance", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to validate remittance account holder: %w", err)
	}

	var holder AccountHolder
	if err := json.Unmarshal(respBody, &holder); err != nil {
		return nil, fmt.Errorf("failed to decode account holder status: %w", err)
	}

	return &holder, nil
}

func (c *Client) GetRemittanceUserInfo(msisdn string) (*BasicUserInfo, error) {
	url := fmt.Sprintf(remittanceUserInfoURL, msisdn)

	respBody, err := c.makeRequest(http.MethodGet, url, "remittance", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get remittance user info: %w", err)
	}

	var userInfo BasicUserInfo
	if err := json.Unmarshal(respBody, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return &userInfo, nil
}
