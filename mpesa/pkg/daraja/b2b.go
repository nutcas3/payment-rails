package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BusinessToBusinessRequestBody struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	SenderIdentifierType   string `json:"SenderIdentifierType"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	Amount                 string `json:"Amount"`
	PartyA                 string `json:"PartyA"`
	PartyB                 string `json:"PartyB"`
	AccountReference       string `json:"AccountReference"`
	Requester              string `json:"Requester"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
}

type BusinessToBusinessResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) BusinessToBusinessPayment(body BusinessToBusinessRequestBody) (*BusinessToBusinessResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, b2bURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make B2B payment request: %w", err)
	}

	// Parse response
	var response BusinessToBusinessResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse B2B payment response: %w", err)
	}

	return &response, nil
}
