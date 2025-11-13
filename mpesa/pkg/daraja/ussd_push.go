package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	UssdPushURL = "/v1/ussdpush/get-msisdn"
)

type UssdPushRequestBody struct {
	PrimaryShortCode  string `json:"primaryShortCode"`
	ReceiverShortCode string `json:"receiverShortCode"`
	Amount            string `json:"amount"`
	PaymentRef        string `json:"paymentRef"`
	CallbackURL       string `json:"callbackUrl"`
	PartnerName       string `json:"partnerName"`
	RequestRefID      string `json:"RequestRefID"`
}

type UssdPushResponse struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

type UssdPushCallbackResponse struct {
	ResultCode       string `json:"resultCode"`
	ResultDesc       string `json:"resultDesc"`
	Amount           string `json:"amount"`
	RequestID        string `json:"requestId"`
	ResultType       string `json:"resultType,omitempty"`
	ConversationID   string `json:"conversationID,omitempty"`
	TransactionID    string `json:"transactionId,omitempty"`
	Status           string `json:"status,omitempty"`
	PaymentReference string `json:"paymentReference,omitempty"`
}

func (s *Service) UssdPush(body UssdPushRequestBody) (*UssdPushResponse, error) {
	if body.PrimaryShortCode == "" {
		return nil, fmt.Errorf("primaryShortCode is required")
	}
	if body.ReceiverShortCode == "" {
		return nil, fmt.Errorf("receiverShortCode is required")
	}
	if body.Amount == "" {
		return nil, fmt.Errorf("amount is required")
	}
	if body.PaymentRef == "" {
		return nil, fmt.Errorf("paymentRef is required")
	}
	if body.CallbackURL == "" {
		return nil, fmt.Errorf("callbackUrl is required")
	}
	if body.PartnerName == "" {
		return nil, fmt.Errorf("partnerName is required")
	}
	if body.RequestRefID == "" {
		return nil, fmt.Errorf("RequestRefID is required")
	}

	respBody, err := s.makeRequest(http.MethodPost, UssdPushURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make USSD Push request: %w", err)
	}

	var response UssdPushResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse USSD Push response: %w", err)
	}

	return &response, nil
}
