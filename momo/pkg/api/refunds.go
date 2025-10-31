package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	refundV1URL       = "/collection/v1_0/refund"
	refundV2URL       = "/collection/v2_0/refund"
	refundStatusURL   = "/collection/v1_0/refund/%s"
)

type RefundRequest struct {
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	ExternalID        string `json:"externalId"`
	PayerMessage      string `json:"payerMessage"`
	PayeeNote         string `json:"payeeNote"`
	ReferenceIDToRefund string `json:"referenceIdToRefund"`
	CallbackURL       string `json:"callbackUrl,omitempty"`
}

type RefundResponse struct {
	ReferenceID string `json:"referenceId"`
	Status      string `json:"status"`
	Reason      string `json:"reason,omitempty"`
}

type RefundStatus struct {
	Amount                  string `json:"amount"`
	Currency                string `json:"currency"`
	FinancialTransactionID  string `json:"financialTransactionId"`
	ExternalID              string `json:"externalId"`
	PayerMessage            string `json:"payerMessage"`
	PayeeNote               string `json:"payeeNote"`
	Status                  string `json:"status"`
	Reason                  string `json:"reason,omitempty"`
}

func (c *Client) Refund(req RefundRequest) (*RefundResponse, error) {
	referenceID := generateUUID()

	headers := map[string]string{
		"X-Reference-Id": referenceID,
	}

	_, err := c.makeRequest(http.MethodPost, refundV1URL, "collection", req, headers)
	if err != nil {
		return nil, fmt.Errorf("refund failed: %w", err)
	}

	return &RefundResponse{
		ReferenceID: referenceID,
		Status:      "PENDING",
	}, nil
}

func (c *Client) RefundV2(req RefundRequest) (*RefundResponse, error) {
	referenceID := generateUUID()

	headers := map[string]string{
		"X-Reference-Id": referenceID,
	}

	_, err := c.makeRequest(http.MethodPost, refundV2URL, "collection", req, headers)
	if err != nil {
		return nil, fmt.Errorf("refund v2 failed: %w", err)
	}

	return &RefundResponse{
		ReferenceID: referenceID,
		Status:      "PENDING",
	}, nil
}

func (c *Client) GetRefundStatus(referenceID string) (*RefundStatus, error) {
	url := fmt.Sprintf(refundStatusURL, referenceID)

	respBody, err := c.makeRequest(http.MethodGet, url, "collection", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get refund status: %w", err)
	}

	var status RefundStatus
	if err := json.Unmarshal(respBody, &status); err != nil {
		return nil, fmt.Errorf("failed to decode refund status: %w", err)
	}

	return &status, nil
}
