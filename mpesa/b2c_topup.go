package mpesa

import (
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

type B2CTopUpRequest struct {
	Initiator          string // The M-Pesa API operator username
	SecurityCredential string // The encrypted password of the M-Pesa API operator
	Amount             string // The transaction amount
	PartyA             string // Your shortcode (from which money will be deducted)
	PartyB             string // The B2C shortcode to which money will be loaded
	AccountReference   string // Optional: Account reference
	Requester          string // Optional: The consumer's mobile number on behalf of whom you are paying
	Remarks            string // Any additional information to be associated with the transaction
	QueueTimeOutURL    string // URL for timeout notifications
	ResultURL          string // URL for result notifications
}

type B2CTopUpResponse struct {
	OriginatorConversationID string
	ConversationID           string
	ResponseCode             string
	ResponseDescription      string
}

func (c *Client) B2CAccountTopUp(req B2CTopUpRequest) (*B2CTopUpResponse, error) {
	internalReq := daraja.B2CTopUpRequest{
		Initiator:              req.Initiator,
		SecurityCredential:     req.SecurityCredential,
		CommandID:              "BusinessPayToBulk", // This API only supports BusinessPayToBulk command
		SenderIdentifierType:   "4",                 // Only type 4 is allowed for this API
		RecieverIdentifierType: "4",                 // Only type 4 is allowed for this API
		Amount:                 req.Amount,
		PartyA:                 req.PartyA,
		PartyB:                 req.PartyB,
		AccountReference:       req.AccountReference,
		Requester:              req.Requester,
		Remarks:                req.Remarks,
		QueueTimeOutURL:        req.QueueTimeOutURL,
		ResultURL:              req.ResultURL,
	}

	resp, err := c.Service.B2CAccountTopUp(internalReq)
	if err != nil {
		return nil, err
	}

	return &B2CTopUpResponse{
		OriginatorConversationID: resp.OriginatorConversationID,
		ConversationID:           resp.ConversationID,
		ResponseCode:             resp.ResponseCode,
		ResponseDescription:      resp.ResponseDescription,
	}, nil
}
