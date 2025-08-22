package api

type CommonResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      any    `json:"data,omitempty"`
}

type AccountBalanceRequest struct {
	CountryCode    string `json:"countryCode"`
	AccountID      string `json:"accountId"`
	AccountType    string `json:"accountType,omitempty"`
	CurrencyCode   string `json:"currencyCode,omitempty"`
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
	CountryCode    string `json:"countryCode"`
	AccountID      string `json:"accountId"`
	AccountType    string `json:"accountType,omitempty"`
	CurrencyCode   string `json:"currencyCode,omitempty"`
}

type Transaction struct {
	TransactionID   string `json:"transactionID"`
	TransactionDate string `json:"transactionDate"`
	ValueDate       string `json:"valueDate"`
	Narration       string `json:"narration"`
	Amount          string `json:"amount"`
	DebitOrCredit   string `json:"debitOrCredit"`
	RunningBalance  string `json:"runningBalance"`
}

type MiniStatementResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      struct {
		Transactions []Transaction `json:"transactions"`
	} `json:"data"`
}

type FullStatementRequest struct {
	CountryCode    string `json:"countryCode"`
	AccountID      string `json:"accountId"`
	AccountType    string `json:"accountType,omitempty"`
	CurrencyCode   string `json:"currencyCode,omitempty"`
	FromDate       string `json:"fromDate"`
	ToDate         string `json:"toDate"`
}

type FullStatementResponse struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
	Data      struct {
		Transactions []Transaction `json:"transactions"`
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
	Type          string `json:"type"`
	Amount        string `json:"amount"`
	CurrencyCode  string `json:"currencyCode"`
	Reference     string `json:"reference"`
	Date          string `json:"date"`
	Description   string `json:"description"`
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
	Source      Source      `json:"source"`
	Destination struct {
		Type          string `json:"type"`
		CountryCode   string `json:"countryCode"`
		Name          string `json:"name"`
		MobileNumber  string `json:"mobileNumber"`
		WalletName    string `json:"walletName"`
	} `json:"destination"`
	Transfer    Transfer    `json:"transfer"`
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
	CountryCode   string `json:"countryCode"`
	CurrencyCode  string `json:"currencyCode"`
	BaseCurrency  string `json:"baseCurrency,omitempty"`
}

// ForexRate represents a forex rate
type ForexRate struct {
	CurrencyCode  string `json:"currencyCode"`
	BuyRate       string `json:"buyRate"`
	SellRate      string `json:"sellRate"`
	MeanRate      string `json:"meanRate"`
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
