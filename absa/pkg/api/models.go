package api

// Account Balance
type AccountBalanceRequest struct {
	AccountNumber string `json:"accountNumber"`
	AccountType   string `json:"accountType,omitempty"`
}

type AccountBalanceResponse struct {
	AccountNumber string `json:"accountNumber"`
	Currency      string `json:"currency"`
	AvailableBalance string `json:"availableBalance"`
	CurrentBalance   string `json:"currentBalance"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}

// Mini Statement
type MiniStatementRequest struct {
	AccountNumber string `json:"accountNumber"`
	StartDate     string `json:"startDate,omitempty"` // YYYY-MM-DD
	EndDate       string `json:"endDate,omitempty"`   // YYYY-MM-DD
	MaxEntries    int    `json:"maxEntries,omitempty"`
}

type StatementEntry struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Type        string `json:"type"` // DEBIT or CREDIT
	Balance     string `json:"balance"`
	Reference   string `json:"reference"`
}

type MiniStatementResponse struct {
	AccountNumber string          `json:"accountNumber"`
	Currency      string          `json:"currency"`
	Entries       []StatementEntry `json:"entries"`
	Status        string          `json:"status"`
	Timestamp     string          `json:"timestamp"`
}

// Full Statement
type FullStatementRequest struct {
	AccountNumber string `json:"accountNumber"`
	StartDate     string `json:"startDate"` // YYYY-MM-DD
	EndDate       string `json:"endDate"`   // YYYY-MM-DD
	Format        string `json:"format,omitempty"` // PDF or CSV
}

type FullStatementResponse struct {
	AccountNumber string `json:"accountNumber"`
	DocumentURL   string `json:"documentUrl,omitempty"`
	DocumentData  string `json:"documentData,omitempty"` // Base64 encoded if not URL
	Format        string `json:"format"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}

// Account Validation
type AccountValidateRequest struct {
	AccountNumber string `json:"accountNumber"`
	BankCode      string `json:"bankCode,omitempty"`
	BranchCode    string `json:"branchCode,omitempty"`
}

type AccountValidateResponse struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	AccountType   string `json:"accountType"`
	BankCode      string `json:"bankCode"`
	BankName      string `json:"bankName"`
	BranchCode    string `json:"branchCode"`
	BranchName    string `json:"branchName"`
	IsActive      bool   `json:"isActive"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}

// Send Money (Bank Transfer)
type SendMoneyRequest struct {
	SourceAccount      string `json:"sourceAccount"`
	DestinationAccount string `json:"destinationAccount"`
	DestinationBankCode string `json:"destinationBankCode,omitempty"`
	Amount             string `json:"amount"`
	Currency           string `json:"currency"`
	Reference          string `json:"reference"`
	Description        string `json:"description,omitempty"`
	CallbackURL        string `json:"callbackUrl,omitempty"`
	BeneficiaryName    string `json:"beneficiaryName,omitempty"`
}

type SendMoneyResponse struct {
	TransactionID string `json:"transactionId"`
	Reference     string `json:"reference"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}

// Mobile Wallet
type MobileWalletRequest struct {
	SourceAccount  string `json:"sourceAccount"`
	MobileNumber   string `json:"mobileNumber"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Reference      string `json:"reference"`
	Description    string `json:"description,omitempty"`
	CallbackURL    string `json:"callbackUrl,omitempty"`
	Provider       string `json:"provider"` // e.g., MPESA, AIRTEL
	CountryCode    string `json:"countryCode,omitempty"`
}

type MobileWalletResponse struct {
	TransactionID string `json:"transactionId"`
	Reference     string `json:"reference"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}

// Bill Payment
type BillPaymentRequest struct {
	SourceAccount string `json:"sourceAccount"`
	BillerCode    string `json:"billerCode"`
	BillerName    string `json:"billerName,omitempty"`
	CustomerReference string `json:"customerReference"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Reference     string `json:"reference"`
	Description   string `json:"description,omitempty"`
	CallbackURL   string `json:"callbackUrl,omitempty"`
}

type BillPaymentResponse struct {
	TransactionID string `json:"transactionId"`
	Reference     string `json:"reference"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
	ReceiptNumber string `json:"receiptNumber,omitempty"`
}

// Receive Money
type ReceiveMoneyRequest struct {
	AccountNumber string `json:"accountNumber"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Reference     string `json:"reference"`
	Description   string `json:"description,omitempty"`
	CallbackURL   string `json:"callbackUrl,omitempty"`
	ExpiryMinutes int    `json:"expiryMinutes,omitempty"`
	PayerName     string `json:"payerName,omitempty"`
	PayerEmail    string `json:"payerEmail,omitempty"`
	PayerPhone    string `json:"payerPhone,omitempty"`
}

type ReceiveMoneyResponse struct {
	PaymentLink   string `json:"paymentLink"`
	PaymentCode   string `json:"paymentCode,omitempty"`
	Reference     string `json:"reference"`
	Status        string `json:"status"`
	ExpiryDate    string `json:"expiryDate,omitempty"`
	Timestamp     string `json:"timestamp"`
}

// Transaction Query
type TransactionQueryRequest struct {
	TransactionID string `json:"transactionId,omitempty"`
	Reference     string `json:"reference,omitempty"`
}

type TransactionQueryResponse struct {
	TransactionID string `json:"transactionId"`
	Reference     string `json:"reference"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Timestamp     string `json:"timestamp"`
	Description   string `json:"description,omitempty"`
	SourceAccount string `json:"sourceAccount,omitempty"`
	Destination   string `json:"destination,omitempty"`
	Fee           string `json:"fee,omitempty"`
}

// Airtime Purchase
type AirtimePurchaseRequest struct {
	SourceAccount string `json:"sourceAccount"`
	MobileNumber  string `json:"mobileNumber"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Reference     string `json:"reference"`
	Provider      string `json:"provider,omitempty"` // e.g., SAFARICOM, AIRTEL
	CountryCode   string `json:"countryCode,omitempty"`
	CallbackURL   string `json:"callbackUrl,omitempty"`
}

type AirtimePurchaseResponse struct {
	TransactionID string `json:"transactionId"`
	Reference     string `json:"reference"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
}
