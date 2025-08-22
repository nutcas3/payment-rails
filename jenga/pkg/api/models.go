package api

import (
	"encoding/json"
)

type CommonResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      any    `json:"data,omitempty"`
}

type AccountBalanceRequest struct {
	CountryCode  string `json:"countryCode"`
	AccountID    string `json:"accountId"`
	AccountType  string `json:"accountType,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
}

type AccountBalanceResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      struct {
		Currency string `json:"currency"`
		Balance  string `json:"balance"`
	} `json:"data"`
}

type MiniStatementRequest struct {
	CountryCode  string `json:"countryCode"`
	AccountID    string `json:"accountId"`
	AccountType  string `json:"accountType,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
}

type MiniStatementTransaction struct {
	ChequeNumber string `json:"chequeNumber"`
	Date         string `json:"date"`
	Description  string `json:"description"`
	Amount       string `json:"amount"`
	Type         string `json:"type"` // "Debit" or "Credit"
}

type MiniStatementResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      struct {
		AccountNumber string                     `json:"accountNumber"`
		Currency      string                     `json:"currency"`
		Balance       json.Number                `json:"balance"`
		Transactions  []MiniStatementTransaction `json:"transactions"`
	} `json:"data"`
}

type FullStatementRequest struct {
	CountryCode  string `json:"countryCode"`
	AccountID    string `json:"accountId"`
	AccountType  string `json:"accountType,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
	FromDate     string `json:"fromDate"`
	ToDate       string `json:"toDate"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

type FullStatementTransaction struct {
	Reference      string `json:"reference"`
	Date           string `json:"date"`
	Amount         string `json:"amount"`
	Serial         string `json:"serial"`
	Description    string `json:"description"`
	PostedDateTime string `json:"postedDateTime"`
	Type           string `json:"type"`
	RunningBalance struct {
		Amount   json.Number `json:"amount"`
		Currency string      `json:"currency"`
	} `json:"runningBalance"`
	TransactionId  string `json:"transactionId"`
}

type FullStatementResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      struct {
		Balance       json.Number             `json:"balance"`
		Currency      string                  `json:"currency"`
		AccountNumber string                  `json:"accountNumber"`
		Transactions  []FullStatementTransaction `json:"transactions"`
	} `json:"data"`
}

type Source struct {
	CountryCode   string `json:"countryCode"`
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
}

// Destination represents the destination account in a money transfer
type Destination struct {
	Type          string `json:"type"`
	CountryCode   string `json:"countryCode"`
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
	// SWIFT specific fields
	BankCode      string `json:"bankCode,omitempty"`
	BranchCode    string `json:"branchCode,omitempty"`
	BankName      string `json:"bankName,omitempty"`
	BankAddress   string `json:"bankAddress,omitempty"`
	SwiftCode     string `json:"swiftCode,omitempty"`
	RoutingNumber string `json:"routingNumber,omitempty"`
}

// Transfer represents the transfer details in a money transfer
type Transfer struct {
	Type         string `json:"type"`
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
	Reference    string `json:"reference"`
	Date         string `json:"date"`
	Description  string `json:"description"`
}

// SendMoneyRequest represents a request to send money
type SendMoneyRequest struct {
	Source      Source      `json:"source"`
	Destination Destination `json:"destination"`
	Transfer    Transfer    `json:"transfer"`
}

// SendMoneyResponse represents the response for send money
type SendMoneyResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
	Data      struct {
		TransactionID string `json:"transactionId"`
		Status        string `json:"status"`
	} `json:"data"`
}

// Mobile Wallet

// MobileWalletRequest represents a request to send money to mobile wallet
type MobileWalletRequest struct {
	Source      Source `json:"source"`
	Destination struct {
		Type         string `json:"type"`
		CountryCode  string `json:"countryCode"`
		Name         string `json:"name"`
		MobileNumber string `json:"mobileNumber"`
		WalletName   string `json:"walletName"`
	} `json:"destination"`
	Transfer struct {
		Type         string `json:"type"`
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
		Reference    string `json:"reference"`
		Date         string `json:"date"`
		Description  string `json:"description"`
		CallbackUrl  string `json:"callbackUrl"`
	} `json:"transfer"`
}

// MobileWalletResponse represents the response for mobile wallet transfer
type MobileWalletResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
	Data      struct {
		TransactionID string `json:"transactionId"`
		Status        string `json:"status"`
	} `json:"data"`
}

// Bill Payments

// BillPaymentRequest represents a request to pay a bill
type BillPaymentRequest struct {
	BillerCode    string `json:"billerCode"`
	AccountNumber string `json:"accountNumber"`
	Amount        string `json:"amount"`
	Reference     string `json:"reference"`
	CurrencyCode  string `json:"currencyCode"`
	Narration     string `json:"narration"`
}

// BillPaymentResponse represents the response for bill payment
type BillPaymentResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
	Data      struct {
		TransactionID string `json:"transactionId"`
		Status        string `json:"status"`
	} `json:"data"`
}

// Airtime

// AirtimePurchaseRequest represents a request to purchase airtime
type AirtimePurchaseRequest struct {
	CustomerMobile string `json:"customerMobile"`
	TelcoCode      string `json:"telcoCode"`
	Amount         string `json:"amount"`
	Reference      string `json:"reference"`
	CurrencyCode   string `json:"currencyCode"`
}

// AirtimePurchaseResponse represents the response for airtime purchase
type AirtimePurchaseResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
	Data      struct {
		TransactionID string `json:"transactionId"`
		Status        string `json:"status"`
	} `json:"data"`
}

// KYC and Identity Verification

// KYCRequest represents a request for KYC verification
type KYCRequest struct {
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	CountryCode    string `json:"countryCode"`
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	DateOfBirth    string `json:"dateOfBirth,omitempty"`
}

// KYCResponse represents the response for KYC verification
type KYCResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      struct {
		FirstName      string `json:"firstName"`
		MiddleName     string `json:"middleName"`
		LastName       string `json:"lastName"`
		FullName       string `json:"fullName"`
		Gender         string `json:"gender"`
		DateOfBirth    string `json:"dateOfBirth"`
		DocumentNumber string `json:"documentNumber"`
		DocumentType   string `json:"documentType"`
		DocumentSerial string `json:"documentSerial"`
		Photo          string `json:"photo"`
	} `json:"data"`
}

// Forex Rates

// ForexRatesRequest represents a request to get forex rates
type ForexRatesRequest struct {
	CountryCode  string `json:"countryCode"`
	CurrencyCode string `json:"currencyCode"`
	BaseCurrency string `json:"baseCurrency,omitempty"`
}

// ForexRate represents a forex rate
type ForexRate struct {
	CurrencyCode string `json:"currencyCode"`
	BuyRate      string `json:"buyRate"`
	SellRate     string `json:"sellRate"`
	MeanRate     string `json:"meanRate"`
}

// ForexRatesResponse represents the response for forex rates
type ForexRatesResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      struct {
		BaseCurrency string      `json:"baseCurrency"`
		Rates        []ForexRate `json:"rates"`
	} `json:"data"`
}
