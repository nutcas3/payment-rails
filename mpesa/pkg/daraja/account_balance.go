package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountBalanceRequestBody struct {
	Initiator          string `json:"Initiator"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	PartyA             int    `json:"PartyA"`
	IdentifierType     int    `json:"IdentifierType"`
	Remarks            string `json:"Remarks"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
}

type AccountBalanceResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) AccountBalance(body AccountBalanceRequestBody) (*AccountBalanceResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, accountBalanceURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make account balance request: %w", err)
	}

	var response AccountBalanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse account balance response: %w", err)
	}

	return &response, nil
}
