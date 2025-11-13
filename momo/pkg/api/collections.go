package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	requestToPayURL        = "/collection/v1_0/requesttopay"
	requestToPayStatusURL  = "/collection/v1_0/requesttopay/%s"
	accountBalanceURL      = "/collection/v1_0/account/balance"
	accountActiveURL       = "/collection/v1_0/accountholder/msisdn/%s/active"
	basicUserInfoURL       = "/collection/v1_0/accountholder/msisdn/%s/basicuserinfo"
)

type RequestToPayRequest struct {
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	ExternalID        string `json:"externalId"`
	Payer             Payer  `json:"payer"`
	PayerMessage      string `json:"payerMessage"`
	PayeeNote         string `json:"payeeNote"`
	CallbackURL       string `json:"callbackUrl,omitempty"`
}

type Payer struct {
	PartyIDType string `json:"partyIdType"`
	PartyID     string `json:"partyId"`
}

type RequestToPayResponse struct {
	ReferenceID string `json:"referenceId"`
	Status      string `json:"status"`
	Reason      string `json:"reason,omitempty"`
}

type TransactionStatus struct {
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	FinancialTransactionID string `json:"financialTransactionId"`
	ExternalID        string `json:"externalId"`
	Payer             Payer  `json:"payer"`
	PayerMessage      string `json:"payerMessage"`
	PayeeNote         string `json:"payeeNote"`
	Status            string `json:"status"`
	Reason            string `json:"reason,omitempty"`
}

type Balance struct {
	AvailableBalance string `json:"availableBalance"`
	Currency         string `json:"currency"`
}

type BasicUserInfo struct {
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Birthdate  string `json:"birthdate"`
	Locale     string `json:"locale"`
	Gender     string `json:"gender"`
}

type AccountHolder struct {
	Result bool `json:"result"`
}

func generateUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (c *Client) RequestToPay(req RequestToPayRequest) (*RequestToPayResponse, error) {
	referenceID := generateUUID()

	headers := map[string]string{
		"X-Reference-Id": referenceID,
	}

	_, err := c.makeRequest(http.MethodPost, requestToPayURL, "collection", req, headers)
	if err != nil {
		return nil, fmt.Errorf("request-to-pay failed: %w", err)
	}

	return &RequestToPayResponse{
		ReferenceID: referenceID,
		Status:      "PENDING",
	}, nil
}

func (c *Client) GetRequestToPayStatus(referenceID string) (*TransactionStatus, error) {
	url := fmt.Sprintf(requestToPayStatusURL, referenceID)

	respBody, err := c.makeRequest(http.MethodGet, url, "collection", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction status: %w", err)
	}

	var status TransactionStatus
	if err := json.Unmarshal(respBody, &status); err != nil {
		return nil, fmt.Errorf("failed to decode transaction status: %w", err)
	}

	return &status, nil
}

func (c *Client) GetAccountBalance() (*Balance, error) {
	respBody, err := c.makeRequest(http.MethodGet, accountBalanceURL, "collection", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account balance: %w", err)
	}

	var balance Balance
	if err := json.Unmarshal(respBody, &balance); err != nil {
		return nil, fmt.Errorf("failed to decode balance: %w", err)
	}

	return &balance, nil
}

func (c *Client) ValidateAccountHolderStatus(msisdn string) (*AccountHolder, error) {
	url := fmt.Sprintf(accountActiveURL, msisdn)

	respBody, err := c.makeRequest(http.MethodGet, url, "collection", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to validate account holder: %w", err)
	}

	var holder AccountHolder
	if err := json.Unmarshal(respBody, &holder); err != nil {
		return nil, fmt.Errorf("failed to decode account holder status: %w", err)
	}

	return &holder, nil
}

func (c *Client) GetBasicUserInfo(msisdn string) (*BasicUserInfo, error) {
	url := fmt.Sprintf(basicUserInfoURL, msisdn)

	respBody, err := c.makeRequest(http.MethodGet, url, "collection", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get basic user info: %w", err)
	}

	var userInfo BasicUserInfo
	if err := json.Unmarshal(respBody, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return &userInfo, nil
}
