package api

import (
	"github.com/shopspring/decimal"
	"time"
)

const (
	StatusSuccess      = "SUCCESS"
	StatusPending      = "PENDING"
	StatusFailed       = "FAILED"
	StatusProcessing   = "PROCESSING"
	StatusCancelled    = "CANCELLED"
	StatusExpired      = "EXPIRED"
)

const (
	TransactionTypeDebit       = "DEBIT"
	TransactionTypeCredit      = "CREDIT"
	TransactionTypeBillPayment = "BILL_PAYMENT"
	TransactionTypeAirtime     = "AIRTIME"
	TransactionTypeTransfer    = "TRANSFER"
)

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
	Type        string          `json:"type"`
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

type FullStatementRequest struct {
	AccountNumber string    `json:"accountNumber"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Format        string    `json:"format,omitempty"`
}

type FullStatementResponse struct {
	AccountNumber string    `json:"accountNumber"`
	DocumentURL   string    `json:"documentUrl,omitempty"`
	DocumentData  string    `json:"documentData,omitempty"`
	Format        string    `json:"format"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}

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

type BulkPaymentItem struct {
	DestinationAccount  string          `json:"destinationAccount"`
	DestinationBankCode string          `json:"destinationBankCode,omitempty"`
	Amount              decimal.Decimal `json:"amount"`
	Reference           string          `json:"reference"`
	Description         string          `json:"description,omitempty"`
	BeneficiaryName     string          `json:"beneficiaryName,omitempty"`
}

type BulkPaymentRequest struct {
	SourceAccount string           `json:"sourceAccount"`
	Currency      string           `json:"currency"`
	CallbackURL   string           `json:"callbackUrl,omitempty"`
	BatchReference string          `json:"batchReference"`
	Items         []BulkPaymentItem `json:"items"`
}

type BulkPaymentItemResponse struct {
	TransactionID string    `json:"transactionId"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"`
	DestinationAccount string `json:"destinationAccount"`
}

type BulkPaymentResponse struct {
	BatchID      string                 `json:"batchId"`
	BatchReference string               `json:"batchReference"`
	Status       string                 `json:"status"`
	Timestamp    time.Time              `json:"timestamp"`
	TotalAmount  decimal.Decimal        `json:"totalAmount"`
	Currency     string                 `json:"currency"`
	ItemCount    int                    `json:"itemCount"`
	SuccessCount int                    `json:"successCount"`
	FailedCount  int                    `json:"failedCount"`
	Items        []BulkPaymentItemResponse `json:"items"`
}

type BulkPaymentStatusRequest struct {
	BatchID string `json:"batchId"`
}

type BulkPaymentStatusResponse struct {
	BatchID      string                 `json:"batchId"`
	BatchReference string               `json:"batchReference"`
	Status       string                 `json:"status"`
	Timestamp    time.Time              `json:"timestamp"`
	TotalAmount  decimal.Decimal        `json:"totalAmount"`
	Currency     string                 `json:"currency"`
	ItemCount    int                    `json:"itemCount"`
	SuccessCount int                    `json:"successCount"`
	FailedCount  int                    `json:"failedCount"`
	Items        []BulkPaymentItemResponse `json:"items"`
}

const (
	FrequencyDaily   = "DAILY"
	FrequencyWeekly  = "WEEKLY"
	FrequencyMonthly = "MONTHLY"
	FrequencyYearly  = "YEARLY"
)

type StandingOrderRequest struct {
	SourceAccount       string          `json:"sourceAccount"`
	DestinationAccount  string          `json:"destinationAccount"`
	DestinationBankCode string          `json:"destinationBankCode,omitempty"`
	Amount              decimal.Decimal `json:"amount"`
	Currency            string          `json:"currency"`
	Reference           string          `json:"reference"`
	Description         string          `json:"description,omitempty"`
	StartDate           time.Time       `json:"startDate"`
	EndDate             time.Time       `json:"endDate,omitempty"`
	Frequency           string          `json:"frequency"` // DAILY, WEEKLY, MONTHLY, YEARLY
	DayOfMonth          int             `json:"dayOfMonth,omitempty"`
	DayOfWeek           int             `json:"dayOfWeek,omitempty"`
	BeneficiaryName     string          `json:"beneficiaryName,omitempty"`
	CallbackURL         string          `json:"callbackUrl,omitempty"`
}

type StandingOrderResponse struct {
	OrderID      string    `json:"orderId"`
	Reference    string    `json:"reference"`
	Status       string    `json:"status"`
	NextRunDate  time.Time `json:"nextRunDate"`
	Timestamp    time.Time `json:"timestamp"`
}

type StandingOrderStatusRequest struct {
	OrderID string `json:"orderId"`
}

type StandingOrderStatusResponse struct {
	OrderID      string          `json:"orderId"`
	Reference    string          `json:"reference"`
	Status       string          `json:"status"`
	Amount       decimal.Decimal `json:"amount"`
	Currency     string          `json:"currency"`
	Frequency    string          `json:"frequency"`
	StartDate    time.Time       `json:"startDate"`
	EndDate      time.Time       `json:"endDate,omitempty"`
	NextRunDate  time.Time       `json:"nextRunDate"`
	LastRunDate  time.Time       `json:"lastRunDate,omitempty"`
	RunCount     int             `json:"runCount"`
	Timestamp    time.Time       `json:"timestamp"`
}

type StandingOrderCancelRequest struct {
	OrderID string `json:"orderId"`
	Reason  string `json:"reason,omitempty"`
}

type StandingOrderCancelResponse struct {
	OrderID   string    `json:"orderId"`
	Reference string    `json:"reference"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type StandingOrderListRequest struct {
	SourceAccount string    `json:"sourceAccount,omitempty"`
	Status        string    `json:"status,omitempty"`
	FromDate      time.Time `json:"fromDate,omitempty"`
	ToDate        time.Time `json:"toDate,omitempty"`
}

type StandingOrderListResponse struct {
	Orders    []StandingOrderStatusResponse `json:"orders"`
	Count     int                           `json:"count"`
	Timestamp time.Time                     `json:"timestamp"`
}

const (
	BeneficiaryTypeBank   = "BANK"
	BeneficiaryTypeMobile = "MOBILE"
	BeneficiaryTypeBiller = "BILLER"
)

type Beneficiary struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	AccountNumber    string    `json:"accountNumber,omitempty"`
	BankCode         string    `json:"bankCode,omitempty"`
	BranchCode       string    `json:"branchCode,omitempty"`
	MobileNumber     string    `json:"mobileNumber,omitempty"`
	MobileProvider   string    `json:"mobileProvider,omitempty"`
	BillerCode       string    `json:"billerCode,omitempty"`
	CustomerReference string    `json:"customerReference,omitempty"`
	Email            string    `json:"email,omitempty"`
	PhoneNumber      string    `json:"phoneNumber,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type BeneficiaryCreateRequest struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	AccountNumber    string `json:"accountNumber,omitempty"`
	BankCode         string `json:"bankCode,omitempty"`
	BranchCode       string `json:"branchCode,omitempty"`
	MobileNumber     string `json:"mobileNumber,omitempty"`
	MobileProvider   string `json:"mobileProvider,omitempty"`
	BillerCode       string `json:"billerCode,omitempty"`
	CustomerReference string `json:"customerReference,omitempty"`
	Email            string `json:"email,omitempty"`
	PhoneNumber      string `json:"phoneNumber,omitempty"`
}

type BeneficiaryCreateResponse struct {
	Beneficiary Beneficiary `json:"beneficiary"`
	Status      string      `json:"status"`
	Timestamp   time.Time   `json:"timestamp"`
}

type BeneficiaryListRequest struct {
	Type string `json:"type,omitempty"` // Filter by type
}

type BeneficiaryListResponse struct {
	Beneficiaries []Beneficiary `json:"beneficiaries"`
	Count         int           `json:"count"`
	Timestamp     time.Time     `json:"timestamp"`
}

type BeneficiaryGetRequest struct {
	ID string `json:"id"`
}

type BeneficiaryGetResponse struct {
	Beneficiary Beneficiary `json:"beneficiary"`
	Timestamp   time.Time   `json:"timestamp"`
}

type BeneficiaryUpdateRequest struct {
	ID               string `json:"id"`
	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	PhoneNumber      string `json:"phoneNumber,omitempty"`
	CustomerReference string `json:"customerReference,omitempty"`
}

type BeneficiaryUpdateResponse struct {
	Beneficiary Beneficiary `json:"beneficiary"`
	Status      string      `json:"status"`
	Timestamp   time.Time   `json:"timestamp"`
}

type BeneficiaryDeleteRequest struct {
	ID string `json:"id"`
}

type BeneficiaryDeleteResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type ForexRateRequest struct {
	SourceCurrency      string          `json:"sourceCurrency"`
	DestinationCurrency string          `json:"destinationCurrency"`
	Amount              decimal.Decimal `json:"amount,omitempty"`
}

type ForexRateResponse struct {
	SourceCurrency      string          `json:"sourceCurrency"`
	DestinationCurrency string          `json:"destinationCurrency"`
	Rate                decimal.Decimal `json:"rate"`
	SourceAmount        decimal.Decimal `json:"sourceAmount,omitempty"`
	DestinationAmount   decimal.Decimal `json:"destinationAmount,omitempty"`
	Timestamp           time.Time       `json:"timestamp"`
}

type ForexTransferRequest struct {
	SourceAccount       string          `json:"sourceAccount"`
	DestinationAccount  string          `json:"destinationAccount"`
	DestinationBankCode string          `json:"destinationBankCode,omitempty"`
	SourceCurrency      string          `json:"sourceCurrency"`
	DestinationCurrency string          `json:"destinationCurrency"`
	SourceAmount        decimal.Decimal `json:"sourceAmount"`
	Reference           string          `json:"reference"`
	Description         string          `json:"description,omitempty"`
	CallbackURL         string          `json:"callbackUrl,omitempty"`
	BeneficiaryName     string          `json:"beneficiaryName,omitempty"`
	BeneficiaryAddress  string          `json:"beneficiaryAddress,omitempty"`
	BeneficiaryCountry  string          `json:"beneficiaryCountry,omitempty"`
	PurposeOfPayment    string          `json:"purposeOfPayment,omitempty"`
}

type ForexTransferResponse struct {
	TransactionID      string          `json:"transactionId"`
	Reference          string          `json:"reference"`
	Status             string          `json:"status"`
	SourceAmount       decimal.Decimal `json:"sourceAmount"`
	DestinationAmount  decimal.Decimal `json:"destinationAmount"`
	ExchangeRate       decimal.Decimal `json:"exchangeRate"`
	Fee                decimal.Decimal `json:"fee,omitempty"`
	EstimatedDelivery  time.Time       `json:"estimatedDelivery,omitempty"`
	Timestamp          time.Time       `json:"timestamp"`
}

type ForexTransferStatusRequest struct {
	TransactionID string `json:"transactionId"`
}

type ForexTransferStatusResponse struct {
	TransactionID      string          `json:"transactionId"`
	Reference          string          `json:"reference"`
	Status             string          `json:"status"`
	SourceAmount       decimal.Decimal `json:"sourceAmount"`
	DestinationAmount  decimal.Decimal `json:"destinationAmount"`
	ExchangeRate       decimal.Decimal `json:"exchangeRate"`
	Fee                decimal.Decimal `json:"fee,omitempty"`
	EstimatedDelivery  time.Time       `json:"estimatedDelivery,omitempty"`
	ActualDelivery     time.Time       `json:"actualDelivery,omitempty"`
	Timestamp          time.Time       `json:"timestamp"`
}

const (
	AuthMethodOTP      = "OTP"
	AuthMethodBiometric = "BIOMETRIC"
	AuthMethod2FA      = "2FA"
)

type OTPRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email,omitempty"`
	Purpose     string `json:"purpose"`
	Reference   string `json:"reference,omitempty"`
}

type OTPResponse struct {
	RequestID  string    `json:"requestId"`
	Status     string    `json:"status"`
	ExpiryTime time.Time `json:"expiryTime"`
	Timestamp  time.Time `json:"timestamp"`
}

type OTPVerifyRequest struct {
	RequestID string `json:"requestId"`
	OTPCode   string `json:"otpCode"`
}

type OTPVerifyResponse struct {
	RequestID string    `json:"requestId"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type TransactionAuthRequest struct {
	TransactionID string `json:"transactionId"`
	AuthMethod    string `json:"authMethod"`
	AuthData      string `json:"authData,omitempty"`
}

type TransactionAuthResponse struct {
	TransactionID string    `json:"transactionId"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}

type DeviceRegistrationRequest struct {
	DeviceID       string `json:"deviceId"`
	DeviceName     string `json:"deviceName"`
	DeviceType     string `json:"deviceType"`
	OperatingSystem string `json:"operatingSystem"`
	AppVersion     string `json:"appVersion"`
}

type DeviceRegistrationResponse struct {
	DeviceID   string    `json:"deviceId"`
	Status     string    `json:"status"`
	Timestamp  time.Time `json:"timestamp"`
}
