package mpesa

import (
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

type BusinessPayBillRequest struct {
	Initiator          string // The M-Pesa API operator username
	SecurityCredential string // The encrypted password of the M-Pesa API operator
	Amount             string // The transaction amount
	PartyA             string // Your shortcode (from which money will be deducted)
	PartyB             string // The shortcode to which money will be moved
	AccountReference   string // The account number to be associated with the payment (up to 13 characters)
	Requester          string // Optional. The consumer's mobile number on behalf of whom you are paying
	Remarks            string // Any additional information to be associated with the transaction
	QueueTimeOutURL    string // URL for timeout notifications
	ResultURL          string // URL for result notifications
	Occasion           string // Optional. Any additional information to be associated with the transaction
}

type BusinessPayBillResponse struct {
	OriginatorConversationID string
	ConversationID           string
	ResponseCode             string
	ResponseDescription      string
}

func (c *Client) BusinessPayBill(req BusinessPayBillRequest) (*BusinessPayBillResponse, error) {
	internalReq := daraja.BusinessPayBillRequest{
		Initiator:              req.Initiator,
		SecurityCredential:     req.SecurityCredential,
		CommandID:              "BusinessPayBill", // This API only supports BusinessPayBill command
		SenderIdentifierType:   "4",               // Only type 4 is allowed for this API
		RecieverIdentifierType: "4",               // Only type 4 is allowed for this API
		Amount:                 req.Amount,
		PartyA:                 req.PartyA,
		PartyB:                 req.PartyB,
		AccountReference:       req.AccountReference,
		Requester:              req.Requester,
		Remarks:                req.Remarks,
		QueueTimeOutURL:        req.QueueTimeOutURL,
		ResultURL:              req.ResultURL,
		Occasion:               req.Occasion,
	}

	resp, err := c.Service.BusinessPayBill(internalReq)
	if err != nil {
		return nil, err
	}

	return &BusinessPayBillResponse{
		OriginatorConversationID: resp.OriginatorConversationID,
		ConversationID:           resp.ConversationID,
		ResponseCode:             resp.ResponseCode,
		ResponseDescription:      resp.ResponseDescription,
	}, nil
}
