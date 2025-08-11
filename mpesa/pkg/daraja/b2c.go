package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type B2CRequestBody struct {
	InitiatorName          string `json:"InitiatorName"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	Amount                 int    `json:"Amount"`
	PartyA                 int    `json:"PartyA"`
	PartyB                 int    `json:"PartyB"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
	Occassion              string `json:"Occassion"`
}

type B2CResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) B2CPayment(body B2CRequestBody) (*B2CResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, b2cURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make B2C payment request: %w", err)
	}

	var response B2CResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse B2C payment response: %w", err)
	}

	return &response, nil
}
