package mpesa

import (
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

type TaxRemittanceRequest struct {
	Initiator          string // The M-Pesa API operator username
	SecurityCredential string // The encrypted password of the M-Pesa API operator
	Amount             string // The transaction amount
	PartyA             string // Your shortcode (from which money will be deducted)
	AccountReference   string // The payment registration number (PRN) issued by KRA
	Remarks            string // Any additional information to be associated with the transaction
	QueueTimeOutURL    string // URL for timeout notifications
	ResultURL          string // URL for result notifications
}

type TaxRemittanceResponse struct {
	OriginatorConversationID string
	ConversationID           string
	ResponseCode             string
	ResponseDescription      string
}

func (c *Client) RemitTax(req TaxRemittanceRequest) (*TaxRemittanceResponse, error) {
	internalReq := daraja.TaxRemittanceRequestBody{
		Initiator:              req.Initiator,
		SecurityCredential:     req.SecurityCredential,
		CommandID:              daraja.CommandIDPayTaxToKRA, // Only PayTaxToKRA is allowed for this API
		SenderIdentifierType:   "4",                         // Only type 4 is allowed for this API
		RecieverIdentifierType: "4",                         // Only type 4 is allowed for this API
		Amount:                 req.Amount,
		PartyA:                 req.PartyA,
		PartyB:                 "572572", // Only 572572 (KRA) is allowed for this API
		AccountReference:       req.AccountReference,
		Remarks:                req.Remarks,
		QueueTimeOutURL:        req.QueueTimeOutURL,
		ResultURL:              req.ResultURL,
	}

	resp, err := c.Service.RemitTax(internalReq)
	if err != nil {
		return nil, err
	}

	return &TaxRemittanceResponse{
		OriginatorConversationID: resp.OriginatorConversationID,
		ConversationID:           resp.ConversationID,
		ResponseCode:             resp.ResponseCode,
		ResponseDescription:      resp.ResponseDescription,
	}, nil
}
