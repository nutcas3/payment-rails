package api

import (
	"github.com/shopspring/decimal"
	"time"
)

// Status constants
const (
	StatusSuccess      = "SUCCESS"
	StatusPending      = "PENDING"
	StatusFailed       = "FAILED"
	StatusProcessing   = "PROCESSING"
	StatusCancelled    = "CANCELLED"
	StatusExpired      = "EXPIRED"
)

// Transaction type constants
const (
	TransactionTypeDebit       = "DEBIT"
	TransactionTypeCredit      = "CREDIT"
	TransactionTypeBillPayment = "BILL_PAYMENT"
	TransactionTypeAirtime     = "AIRTIME"
	TransactionTypeTransfer    = "TRANSFER"
)

// Account Balance
type AccountBalanceRequest struct {
	AccountNumber string `json:"accountNumber"`
	AccountType   string `json:"accountType,omitempty"`
}

type AccountBalanceResponse struct {
	AccountNumber    string          `json:"accountNumber"`
	Currency         string          `json:"currency"`
	AvailableBalance decimal.Decimal `json:"availableBalance"`
	CurrentBalance   decimal.Decimal `json:"currentBalance"`
	Status           string          `json:"status"`
	Timestamp        time.Time       `json:"timestamp"`
}

// Mini Statement
type MiniStatementRequest struct {
	AccountNumber string    `json:"accountNumber"`
	StartDate     time.Time `json:"startDate,omitempty"`
	EndDate       time.Time `json:"endDate,omitempty"`
	MaxEntries    int       `json:"maxEntries,omitempty"`
}

type StatementEntry struct {
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
	Amount      decimal.Decimal `json:"amount"`
	Type        string          `json:"type"` // DEBIT or CREDIT
	Balance     decimal.Decimal `json:"balance"`
	Reference   string          `json:"reference"`
}

type MiniStatementResponse struct {
	AccountNumber string           `json:"accountNumber"`
	Currency      string           `json:"currency"`
	Entries       []StatementEntry `json:"entries"`
	Status        string           `json:"status"`
	Timestamp     time.Time        `json:"timestamp"`
}

// Full Statement
type FullStatementRequest struct {
	AccountNumber string    `json:"accountNumber"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Format        string    `json:"format,omitempty"` // PDF or CSV
}

type FullStatementResponse struct {
	AccountNumber string    `json:"accountNumber"`
	DocumentURL   string    `json:"documentUrl,omitempty"`
	DocumentData  string    `json:"documentData,omitempty"` // Base64 encoded if not URL
	Format        string    `json:"format"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}

// Account Validation
type AccountValidateRequest struct {
	AccountNumber string `json:"accountNumber"`
	BankCode      string `json:"bankCode,omitempty"`
	BranchCode    string `json:"branchCode,omitempty"`
}

type AccountValidateResponse struct {
	AccountNumber string    `json:"accountNumber"`
	AccountName   string    `json:"accountName"`
	AccountType   string    `json:"accountType"`
	BankCode      string    `json:"bankCode"`
	BankName      string    `json:"bankName"`
	BranchCode    string    `json:"branchCode"`
	BranchName    string    `json:"branchName"`
	IsActive      bool      `json:"isActive"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}

// Send Money (Bank Transfer)
type SendMoneyRequest struct {
	SourceAccount       string          `json:"sourceAccount"`
	DestinationAccount  string          `json:"destinationAccount"`
	DestinationBankCode string          `json:"destinationBankCode,omitempty"`
	Amount              decimal.Decimal `json:"amount"`
	Currency            string          `json:"currency"`
	Reference           string          `json:"reference"`
	Description         string          `json:"description,omitempty"`
	CallbackURL         string          `json:"callbackUrl,omitempty"`
	BeneficiaryName     string          `json:"beneficiaryName,omitempty"`
}

type SendMoneyResponse struct {
	TransactionID string    `json:"transactionId"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}

// Mobile Wallet
type MobileWalletRequest struct {
	SourceAccount string          `json:"sourceAccount"`
	MobileNumber  string          `json:"mobileNumber"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Reference     string          `json:"reference"`
	Description   string          `json:"description,omitempty"`
	CallbackURL   string          `json:"callbackUrl,omitempty"`
	Provider      string          `json:"provider"` // e.g., MPESA, AIRTEL
	CountryCode   string          `json:"countryCode,omitempty"`
}

type MobileWalletResponse struct {
	TransactionID string    `json:"transactionId"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}

// Bill Payment
type BillPaymentRequest struct {
	SourceAccount     string          `json:"sourceAccount"`
	BillerCode        string          `json:"billerCode"`
	BillerName        string          `json:"billerName,omitempty"`
	CustomerReference string          `json:"customerReference"`
	Amount            decimal.Decimal `json:"amount"`
	Currency          string          `json:"currency"`
	Reference         string          `json:"reference"`
	Description       string          `json:"description,omitempty"`
	CallbackURL       string          `json:"callbackUrl,omitempty"`
}

type BillPaymentResponse struct {
	TransactionID string    `json:"transactionId"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
	ReceiptNumber string    `json:"receiptNumber,omitempty"`
}

// Receive Money
type ReceiveMoneyRequest struct {
	AccountNumber string          `json:"accountNumber"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Reference     string          `json:"reference"`
	Description   string          `json:"description,omitempty"`
	CallbackURL   string          `json:"callbackUrl,omitempty"`
	ExpiryMinutes int             `json:"expiryMinutes,omitempty"`
	PayerName     string          `json:"payerName,omitempty"`
	PayerEmail    string          `json:"payerEmail,omitempty"`
	PayerPhone    string          `json:"payerPhone,omitempty"`
}

type ReceiveMoneyResponse struct {
	PaymentLink string    `json:"paymentLink"`
	PaymentCode string    `json:"paymentCode,omitempty"`
	Reference   string    `json:"reference"`
	Status      string    `json:"status"`
	ExpiryDate  time.Time `json:"expiryDate,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

// Transaction Query
type TransactionQueryRequest struct {
	TransactionID string    `json:"transactionId,omitempty"`
	Reference     string    `json:"reference,omitempty"`
	FromDate      time.Time `json:"fromDate,omitempty"`
	ToDate        time.Time `json:"toDate,omitempty"`
}

type TransactionQueryResponse struct {
	TransactionID string          `json:"transactionId"`
	Reference     string          `json:"reference"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Status        string          `json:"status"`
	Type          string          `json:"type"`
	Timestamp     time.Time       `json:"timestamp"`
	Description   string          `json:"description,omitempty"`
	SourceAccount string          `json:"sourceAccount,omitempty"`
	Destination   string          `json:"destination,omitempty"`
	Fee           decimal.Decimal `json:"fee,omitempty"`
}

// Airtime Purchase
type AirtimePurchaseRequest struct {
	SourceAccount string          `json:"sourceAccount"`
	MobileNumber  string          `json:"mobileNumber"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Reference     string          `json:"reference"`
	Provider      string          `json:"provider,omitempty"` // e.g., SAFARICOM, AIRTEL
	CountryCode   string          `json:"countryCode,omitempty"`
	CallbackURL   string          `json:"callbackUrl,omitempty"`
}

type AirtimePurchaseResponse struct {
	TransactionID string    `json:"transactionId"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}
