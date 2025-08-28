package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ratibaURL = "/standingorder/v1/createStandingOrderExternal"

type RatibaRequest struct {
	StandingOrderName           string `json:"StandingOrderName"`
	StartDate                   string `json:"StartDate"`
	EndDate                     string `json:"EndDate"`
	BusinessShortCode           string `json:"BusinessShortCode"`
	TransactionType             string `json:"TransactionType"`
	ReceiverPartyIdentifierType string `json:"ReceiverPartyIdentifierType"`
	Amount                      string `json:"Amount"`
	PartyA                      string `json:"PartyA"`
	CallBackURL                 string `json:"CallBackURL"`
	AccountReference            string `json:"AccountReference"`
	TransactionDesc             string `json:"TransactionDesc"`
	Frequency                   string `json:"Frequency"`
}

type RatibaResponseHeader struct {
	ResponseRefID       string `json:"responseRefID"`
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
	ResultDesc          string `json:"ResultDesc,omitempty"`
}

type RatibaResponseBody struct {
	ResponseDescription string `json:"responseDescription"`
	ResponseCode        string `json:"responseCode"`
}

type RatibaResponse struct {
	ResponseHeader RatibaResponseHeader `json:"ResponseHeader"`
	ResponseBody   RatibaResponseBody   `json:"ResponseBody"`
}

type RatibaCallbackDataItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RatibaCallbackBody struct {
	ResponseData []RatibaCallbackDataItem `json:"responseData"`
}

type RatibaCallback struct {
	ResponseHeader RatibaResponseHeader `json:"responseHeader"`
	ResponseBody   RatibaCallbackBody   `json:"responseBody"`
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
	TransactionTypeMerchant = "Standing Order Customer Pay Marchant" // For Buy Goods
)

const (
	ReceiverTypeTill    = "2" // For Till Number (Merchant)
	ReceiverTypePaybill = "4" // For Business Short Code (PayBill)
)

func (s *Service) CreateStandingOrder(req RatibaRequest) (*RatibaResponse, error) {
	if req.StandingOrderName == "" {
		return nil, fmt.Errorf("StandingOrderName is required")
	}
	if req.StartDate == "" {
		return nil, fmt.Errorf("StartDate is required")
	}
	if req.EndDate == "" {
		return nil, fmt.Errorf("end date is required")
	}
	if req.BusinessShortCode == "" {
		return nil, fmt.Errorf("business shortcode is required")
	}
	if req.TransactionType == "" {
		return nil, fmt.Errorf("transaction type is required")
	}
	if req.Amount == "" {
		return nil, fmt.Errorf("amount is required")
	}
	if req.PartyA == "" {
		return nil, fmt.Errorf("party A (phone number) is required")
	}
	if req.CallBackURL == "" {
		return nil, fmt.Errorf("callback URL is required")
	}
	if req.Frequency == "" {
		return nil, fmt.Errorf("frequency is required")
	}

	if req.ReceiverPartyIdentifierType == "" {
		req.ReceiverPartyIdentifierType = ReceiverTypePaybill
	}

	respBody, err := s.makeRequest(http.MethodPost, ratibaURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create standing order: %w", err)
	}

	var response RatibaResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse standing order response: %w", err)
	}

	return &response, nil
}
