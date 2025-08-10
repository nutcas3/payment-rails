package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ReversalRequestBody struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	TransactionID          string `json:"TransactionID"`
	Amount                 int    `json:"Amount"`
	ReceiverParty          int    `json:"ReceiverParty"`
	RecieverIdentifierType int    `json:"RecieverIdentifierType"`
	ResultURL              string `json:"ResultURL"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	Remarks                string `json:"Remarks"`
	Occasion               string `json:"Occasion"`
}

type ReversalResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) Reversal(body ReversalRequestBody) (*ReversalResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, reversalURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make reversal request: %w", err)
	}

	var response ReversalResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse reversal response: %w", err)
	}

	return &response, nil
}
