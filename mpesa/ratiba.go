package mpesa

import (
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

type RatibaRequest struct {
	StandingOrderName string
	StartDate         string
	EndDate           string
	BusinessShortCode string
	TransactionType   string
	IdentifierType    string
	Amount            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
	Frequency         string
}

type RatibaResponse struct {
	ResponseRefID       string
	ResponseCode        string
	ResponseDescription string
	ResultDesc          string
}

const (
	FrequencyOneOff     = "1" // One Off
	FrequencyDaily      = "2" // Daily
	FrequencyWeekly     = "3" // Weekly
	FrequencyMonthly    = "4" // Monthly
	FrequencyBiMonthly  = "5" // Bi-Monthly
	FrequencyQuarterly  = "6" // Quarterly
	FrequencyHalfYearly = "7" // Half Year
	FrequencyYearly     = "8" // Yearly
)

const (
	TransactionTypePayBill  = "Standing Order Customer Pay Bill"     // For Paybill
	TransactionTypeMerchant = "Standing Order Customer Pay Merchant" // For Buy Goods
)

const (
	ReceiverTypeTill    = "2" // For Till Number (Merchant)
	ReceiverTypePaybill = "4" // For Business Short Code (PayBill)
)

func (c *Client) CreateStandingOrder(req RatibaRequest) (*RatibaResponse, error) {
	internalReq := daraja.RatibaRequest{
		StandingOrderName:           req.StandingOrderName,
		StartDate:                   req.StartDate,
		EndDate:                     req.EndDate,
		BusinessShortCode:           req.BusinessShortCode,
		TransactionType:             req.TransactionType,
		ReceiverPartyIdentifierType: req.IdentifierType,
		Amount:                      req.Amount,
		PartyA:                      req.PhoneNumber,
		CallBackURL:                 req.CallBackURL,
		AccountReference:            req.AccountReference,
		TransactionDesc:             req.TransactionDesc,
		Frequency:                   req.Frequency,
	}

	resp, err := c.Service.CreateStandingOrder(internalReq)
	if err != nil {
		return nil, err
	}

	return &RatibaResponse{
		ResponseRefID:       resp.ResponseHeader.ResponseRefID,
		ResponseCode:        resp.ResponseHeader.ResponseCode,
		ResponseDescription: resp.ResponseHeader.ResponseDescription,
		ResultDesc:          resp.ResponseHeader.ResultDesc,
	}, nil
}
