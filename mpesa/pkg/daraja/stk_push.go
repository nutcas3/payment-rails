package daraja

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type STKPushBody struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

type STKPushResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}

type STKPushQueryBody struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}

type STKPushQueryResponse struct {
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResultCode          string `json:"ResultCode"`
	ResultDesc          string `json:"ResultDesc"`
}

func (s *Service) InitiateStkPush(body STKPushBody) (*STKPushResponse, error) {
	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(body.BusinessShortCode + s.passKey + timestamp))

	payload := STKPushBody{
		BusinessShortCode: body.BusinessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   body.TransactionType,
		Amount:            body.Amount,
		PartyA:            body.PartyA,
		PartyB:            body.PartyB,
		PhoneNumber:       body.PhoneNumber,
		CallBackURL:       body.CallBackURL,
		AccountReference:  body.AccountReference,
		TransactionDesc:   body.TransactionDesc,
	}

	respBody, err := s.makeRequest(http.MethodPost, stkPushURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to make STK push request: %w", err)
	}
	var response STKPushResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse STK push response: %w", err)
	}

	return &response, nil
}

func (s *Service) QueryStkPush(businessShortCode, checkoutRequestID string) (*STKPushQueryResponse, error) {
	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(businessShortCode + s.passKey + timestamp))

	payload := STKPushQueryBody{
		BusinessShortCode: businessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		CheckoutRequestID: checkoutRequestID,
	}

	respBody, err := s.makeRequest(http.MethodPost, stkPushQueryURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to make STK push query request: %w", err)
	}

	var response STKPushQueryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse STK push query response: %w", err)
	}

	return &response, nil
}
