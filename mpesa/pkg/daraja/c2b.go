package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterC2BURLBody struct {
	ShortCode       string `json:"ShortCode"`
	ResponseType    string `json:"ResponseType"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ValidationURL   string `json:"ValidationURL"`
}

type RegisterC2BURLResponse struct {
	OriginatorCoversationID string `json:"OriginatorCoversationID"`
	ResponseCode            string `json:"ResponseCode"`
	ResponseDescription     string `json:"ResponseDescription"`
}

type C2BSimulateRequestBody struct {
	ShortCode     int    `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        int    `json:"Amount"`
	Msisdn        int    `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}

type C2BSimulateResponse struct {
	OriginatorCoversationID string `json:"OriginatorCoversationID"`
	ResponseCode            string `json:"ResponseCode"`
	ResponseDescription     string `json:"ResponseDescription"`
}

func (s *Service) C2BRegisterURL(body RegisterC2BURLBody) (*RegisterC2BURLResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, c2bRegisterURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make C2B register URL request: %w", err)
	}

	var response RegisterC2BURLResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse C2B register URL response: %w", err)
	}

	return &response, nil
}

func (s *Service) C2BSimulate(body C2BSimulateRequestBody) (*C2BSimulateResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, c2bSimulateURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make C2B simulate request: %w", err)
	}

	var response C2BSimulateResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse C2B simulate response: %w", err)
	}

	return &response, nil
}
