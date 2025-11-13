package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	TaxRemittanceURL = "/mpesa/b2b/v1/remittax"

	CommandIDPayTaxToKRA = "PayTaxToKRA"
)

type TaxRemittanceRequestBody struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	SenderIdentifierType   string `json:"SenderIdentifierType"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	Amount                 string `json:"Amount"`
	PartyA                 string `json:"PartyA"`
	PartyB                 string `json:"PartyB"`
	AccountReference       string `json:"AccountReference"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
}

type TaxRemittanceResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (s *Service) RemitTax(body TaxRemittanceRequestBody) (*TaxRemittanceResponse, error) { // Validate required fields
	if body.Initiator == "" {
		return nil, fmt.Errorf("initiator is required")
	}
	if body.SecurityCredential == "" {
		return nil, fmt.Errorf("security credential is required")
	}
	if body.CommandID != CommandIDPayTaxToKRA {
		return nil, fmt.Errorf("invalid CommandID: only PayTaxToKRA is allowed for this API")
	}
	if body.SenderIdentifierType != "4" {
		return nil, fmt.Errorf("invalid SenderIdentifierType: only 4 is allowed for this API")
	}
	if body.RecieverIdentifierType != "4" {
		return nil, fmt.Errorf("invalid RecieverIdentifierType: only 4 is allowed for this API")
	}
	if body.PartyA == "" {
		return nil, fmt.Errorf("partyA is required")
	}
	if body.PartyB != "572572" {
		return nil, fmt.Errorf("invalid PartyB: only 572572 (KRA) is allowed for this API")
	}
	if body.AccountReference == "" {
		return nil, fmt.Errorf("account reference (PRN) is required")
	}
	if body.QueueTimeOutURL == "" {
		return nil, fmt.Errorf("queueTimeOutURL is required")
	}
	if body.ResultURL == "" {
		return nil, fmt.Errorf("resultURL is required")
	}

	respBody, err := s.makeRequest(http.MethodPost, TaxRemittanceURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to make Tax Remittance request: %w", err)
	}

	var response TaxRemittanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse Tax Remittance response: %w", err)
	}

	return &response, nil
}
