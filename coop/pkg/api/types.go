package api

import "time"

type BaseRequest struct {
	MessageReference string `json:"messageReference"`
}

type BaseResponse struct {
	MessageReference string `json:"messageReference"`
	ResponseCode     string `json:"responseCode"`
	ResponseMessage  string `json:"responseMessage"`
}

type AccountBalanceRequest struct {
	BaseRequest
	AccountNumber string `json:"accountNumber"`
}

type AccountBalanceResponse struct {
	BaseResponse
	AccountNumber    string  `json:"accountNumber"`
	AccountName      string  `json:"accountName"`
	Currency         string  `json:"currency"`
	ProductName      string  `json:"productName"`
	ClearedBalance   float64 `json:"clearedBalance"`
	BookedBalance    float64 `json:"bookedBalance"`
	UnclearedBalance float64 `json:"unclearedBalance"`
	BlockedBalance   float64 `json:"blockedBalance"`
}

type AccountTransactionsRequest struct {
	BaseRequest
	AccountNumber      string `json:"accountNumber"`
	NoOfTransactions   string `json:"noOfTransactions"`
}

type Transaction struct {
	TransactionID       string    `json:"transactionId"`
	TransactionDate     time.Time `json:"transactionDate"`
	ValueDate           time.Time `json:"valueDate"`
	Narration           string    `json:"narration"`
	TransactionType     string    `json:"transactionType"`
	ServicePoint        string    `json:"servicePoint"`
	AccountNumber       string    `json:"accountNumber"`
	Currency            string    `json:"currency"`
	Amount              float64   `json:"amount"`
	SerialNumber        string    `json:"serialNumber"`
	DebitCreditIndicator string   `json:"debitCreditIndicator"`
	RunningClearedBalance float64 `json:"runningClearedBalance"`
	RunningBookBalance   float64   `json:"runningBookBalance"`
}

type AccountTransactionsResponse struct {
	BaseResponse
	AccountNumber string        `json:"accountNumber"`
	AccountName   string        `json:"accountName"`
	Currency      string        `json:"currency"`
	ProductName   string        `json:"productName"`
	Transactions  []Transaction `json:"transactions"`
}

type ExchangeRateRequest struct {
	BaseRequest
	FromCurrencyCode string `json:"fromCurrencyCode"`
	ToCurrencyCode   string `json:"toCurrencyCode"`
}

type ExchangeRateResponse struct {
	BaseResponse
	FromCurrencyCode string  `json:"fromCurrencyCode"`
	ToCurrencyCode   string  `json:"toCurrencyCode"`
	RateType         string  `json:"rateType"`
	Rate             float64 `json:"rate"`
	Tolerance        float64 `json:"tolerance"`
	MultiplyDivide   string  `json:"multiplyDivide"`
}

type Destination struct {
	ReferenceNumber   string  `json:"referenceNumber"`
	AccountNumber     string  `json:"accountNumber"`
	Amount            float64 `json:"amount"`
	TransactionCurrency string `json:"transactionCurrency"`
	Narration         string  `json:"narration"`
}

type IFTRequest struct {
	BaseRequest
	AccountNumber       string        `json:"accountNumber"`
	Amount              float64       `json:"amount"`
	TransactionCurrency string        `json:"transactionCurrency"`
	Narration           string        `json:"narration"`
	Destinations        []Destination `json:"destinations"`
}

type IFTResponse struct {
	BaseResponse
	TransactionID string `json:"transactionId"`
}

type PesaLinkDestination struct {
	ReferenceNumber     string  `json:"referenceNumber"`
	DestinationBank     string  `json:"destinationBank"`
	AccountNumber       string  `json:"accountNumber"`
	Amount              float64 `json:"amount"`
	TransactionCurrency string  `json:"transactionCurrency"`
	Narration           string  `json:"narration"`
}

type PesaLinkRequest struct {
	BaseRequest
	AccountNumber       string                `json:"accountNumber"`
	Amount              float64               `json:"amount"`
	TransactionCurrency string                `json:"transactionCurrency"`
	Narration           string                `json:"narration"`
	Destinations        []PesaLinkDestination `json:"destinations"`
}

type PesaLinkResponse struct {
	BaseResponse
	TransactionID string `json:"transactionId"`
}

type TransactionStatusRequest struct {
	BaseRequest
}

type TransactionStatusResponse struct {
	BaseResponse
	TransactionID     string `json:"transactionId"`
	TransactionStatus string `json:"transactionStatus"`
	TransactionAmount float64 `json:"transactionAmount"`
	TransactionDate   time.Time `json:"transactionDate"`
}
