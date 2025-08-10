package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TransactionStatusRequestBody struct {
	Initiator            string `json:"Initiator"`
	SecurityCredential   string `json:"SecurityCredential"`
	CommandID            string `json:"CommandID"`
	TransactionID        string `json:"TransactionID"`
	PartyA               int    `json:"PartyA"`
	IdentifierType       int    `json:"IdentifierType"`
	ResultURL            string `json:"ResultURL"`
	QueueTimeOutURL      string `json:"QueueTimeOutURL"`
	Remarks              string `json:"Remarks"`
	Occassion            string `json:"Occassion"`
}

type TransactionStatusResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) TransactionStatus(body TransactionStatusRequestBody) (*TransactionStatusResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, transactionStatusURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make transaction status request: %w", err)
	}

	var response TransactionStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse transaction status response: %w", err)
	}

	return &response, nil
}
