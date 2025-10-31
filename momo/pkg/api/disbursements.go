package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	transferURL               = "/disbursement/v1_0/transfer"
	transferStatusURL         = "/disbursement/v1_0/transfer/%s"
	disbursementBalanceURL    = "/disbursement/v1_0/account/balance"
	disbursementActiveURL     = "/disbursement/v1_0/accountholder/msisdn/%s/active"
	disbursementUserInfoURL   = "/disbursement/v1_0/accountholder/msisdn/%s/basicuserinfo"
)

type TransferRequest struct {
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	ExternalID        string `json:"externalId"`
	Payee             Payee  `json:"payee"`
	PayerMessage      string `json:"payerMessage"`
	PayeeNote         string `json:"payeeNote"`
	CallbackURL       string `json:"callbackUrl,omitempty"`
}

type Payee struct {
	PartyIDType string `json:"partyIdType"`
	PartyID     string `json:"partyId"`
}

type TransferResponse struct {
	ReferenceID string `json:"referenceId"`
	Status      string `json:"status"`
	Reason      string `json:"reason,omitempty"`
}

func (c *Client) Transfer(req TransferRequest) (*TransferResponse, error) {
	referenceID := generateUUID()

	headers := map[string]string{
		"X-Reference-Id": referenceID,
	}

	_, err := c.makeRequest(http.MethodPost, transferURL, "disbursement", req, headers)
	if err != nil {
		return nil, fmt.Errorf("transfer failed: %w", err)
	}

	return &TransferResponse{
		ReferenceID: referenceID,
		Status:      "PENDING",
	}, nil
}

func (c *Client) GetTransferStatus(referenceID string) (*TransactionStatus, error) {
	url := fmt.Sprintf(transferStatusURL, referenceID)

	respBody, err := c.makeRequest(http.MethodGet, url, "disbursement", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get transfer status: %w", err)
	}

	var status TransactionStatus
	if err := json.Unmarshal(respBody, &status); err != nil {
		return nil, fmt.Errorf("failed to decode transfer status: %w", err)
	}

	return &status, nil
}

func (c *Client) GetDisbursementBalance() (*Balance, error) {
	respBody, err := c.makeRequest(http.MethodGet, disbursementBalanceURL, "disbursement", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get disbursement balance: %w", err)
	}

	var balance Balance
	if err := json.Unmarshal(respBody, &balance); err != nil {
		return nil, fmt.Errorf("failed to decode balance: %w", err)
	}

	return &balance, nil
}

func (c *Client) ValidateDisbursementAccountHolder(msisdn string) (*AccountHolder, error) {
	url := fmt.Sprintf(disbursementActiveURL, msisdn)

	respBody, err := c.makeRequest(http.MethodGet, url, "disbursement", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to validate disbursement account holder: %w", err)
	}

	var holder AccountHolder
	if err := json.Unmarshal(respBody, &holder); err != nil {
		return nil, fmt.Errorf("failed to decode account holder status: %w", err)
	}

	return &holder, nil
}

func (c *Client) GetDisbursementUserInfo(msisdn string) (*BasicUserInfo, error) {
	url := fmt.Sprintf(disbursementUserInfoURL, msisdn)

	respBody, err := c.makeRequest(http.MethodGet, url, "disbursement", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get disbursement user info: %w", err)
	}

	var userInfo BasicUserInfo
	if err := json.Unmarshal(respBody, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return &userInfo, nil
}
