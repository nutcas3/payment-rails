package mpesa

import (
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

type UssdPushRequest struct {
	PrimaryShortCode  string // The merchant's till (organization sending money) shortCode/tillNumber
	ReceiverShortCode string // The vendor(payBill Account) receiving the amount from the merchant
	Amount            string // Amount to be sent to vendor
	PaymentRef        string // Reference to the payment being made
	CallbackURL       string // Endpoint to send back the confirmation response
	PartnerName       string // Organization friendly name used by the vendor
	RequestRefID      string // Random unique identifier for tracking the process
}

type UssdPushResponse struct {
	Code   string // Shows if the push was successful (0) or failed
	Status string // Status message
}

type UssdPushCallbackResponse struct {
	ResultCode       string // Result code (0 for success, other codes for failure)
	ResultDesc       string // Result description
	Amount           string // Transaction amount
	RequestID        string // Unique identifier of the request
	ResultType       string // Status code (usually 0)
	ConversationID   string // Global unique identifier for the transaction
	TransactionID    string // Mpesa Receipt Number (only for successful transactions)
	Status           string // Status of the transaction (SUCCESS/FAILED)
	PaymentReference string // Reference to the payment
}

func (c *Client) UssdPush(req UssdPushRequest) (*UssdPushResponse, error) {
	internalReq := daraja.UssdPushRequestBody{
		PrimaryShortCode:  req.PrimaryShortCode,
		ReceiverShortCode: req.ReceiverShortCode,
		Amount:            req.Amount,
		PaymentRef:        req.PaymentRef,
		CallbackURL:       req.CallbackURL,
		PartnerName:       req.PartnerName,
		RequestRefID:      req.RequestRefID,
	}

	resp, err := c.Service.UssdPush(internalReq)
	if err != nil {
		return nil, err
	}

	return &UssdPushResponse{
		Code:   resp.Code,
		Status: resp.Status,
	}, nil
}
