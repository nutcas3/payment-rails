package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BusinessPayBillRequest struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	SenderIdentifierType   string `json:"SenderIdentifierType"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	Amount                 string `json:"Amount"`
	PartyA                 string `json:"PartyA"`
	PartyB                 string `json:"PartyB"`
	AccountReference       string `json:"AccountReference"`
	Requester              string `json:"Requester,omitempty"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
	Occasion               string `json:"Occasion,omitempty"`
}

type BusinessPayBillResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) BusinessPayBill(req BusinessPayBillRequest) (*BusinessPayBillResponse, error) {
	if req.CommandID != "BusinessPayBill" {
		return nil, fmt.Errorf("invalid CommandID: only BusinessPayBill is allowed for this API")
	}

	if req.SenderIdentifierType != "4" {
		return nil, fmt.Errorf("invalid SenderIdentifierType: only 4 is allowed for this API")
	}

	if req.RecieverIdentifierType != "4" {
		return nil, fmt.Errorf("invalid RecieverIdentifierType: only 4 is allowed for this API")
	}

	respBody, err := s.makeRequest(http.MethodPost, b2bURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to make Business Pay Bill request: %w", err)
	}

	var response BusinessPayBillResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse Business Pay Bill response: %w", err)
	}

	return &response, nil
}
