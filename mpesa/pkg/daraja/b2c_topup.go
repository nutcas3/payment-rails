package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type B2CTopUpRequest struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	SenderIdentifierType   string `json:"SenderIdentifierType"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	Amount                 string `json:"Amount"`
	PartyA                 string `json:"PartyA"`
	PartyB                 string `json:"PartyB"`
	AccountReference       string `json:"AccountReference,omitempty"`
	Requester              string `json:"Requester,omitempty"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
}

type B2CTopUpResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) B2CAccountTopUp(req B2CTopUpRequest) (*B2CTopUpResponse, error) {
	if req.CommandID != "BusinessPayToBulk" {
		return nil, fmt.Errorf("invalid CommandID: only BusinessPayToBulk is allowed for this API")
	}

	if req.SenderIdentifierType != "4" {
		return nil, fmt.Errorf("invalid SenderIdentifierType: only 4 is allowed for this API")
	}

	if req.RecieverIdentifierType != "4" {
		return nil, fmt.Errorf("invalid RecieverIdentifierType: only 4 is allowed for this API")
	}

	respBody, err := s.makeRequest(http.MethodPost, b2bURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to make B2C Account Top Up request: %w", err)
	}

	var response B2CTopUpResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse B2C Account Top Up response: %w", err)
	}

	return &response, nil
}
