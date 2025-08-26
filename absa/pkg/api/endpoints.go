package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// Account Balance
func (c *Client) GetAccountBalance(req AccountBalanceRequest) (*AccountBalanceResponse, error) {
	endpoint := "/accounts/balance"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response AccountBalanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing account balance response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Mini Statement
func (c *Client) GetMiniStatement(req MiniStatementRequest) (*MiniStatementResponse, error) {
	endpoint := "/accounts/mini-statement"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response MiniStatementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing mini statement response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Full Statement
func (c *Client) GetFullStatement(req FullStatementRequest) (*FullStatementResponse, error) {
	endpoint := "/accounts/full-statement"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response FullStatementResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing full statement response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Account Validation
func (c *Client) ValidateAccount(req AccountValidateRequest) (*AccountValidateResponse, error) {
	endpoint := "/accounts/validate"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response AccountValidateResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing account validation response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Send Money (Bank Transfer)
func (c *Client) SendMoney(req SendMoneyRequest) (*SendMoneyResponse, error) {
	endpoint := "/payments/bank-transfer"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response SendMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing send money response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Internal Bank Transfer
func (c *Client) SendInternalBankTransfer(req SendMoneyRequest) (*SendMoneyResponse, error) {
	endpoint := "/payments/internal-transfer"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response SendMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing internal transfer response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Mobile Wallet
func (c *Client) SendToMobileWallet(req MobileWalletRequest) (*MobileWalletResponse, error) {
	endpoint := "/payments/mobile-wallet"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response MobileWalletResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing mobile wallet response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Bill Payment
func (c *Client) PayBill(req BillPaymentRequest) (*BillPaymentResponse, error) {
	endpoint := "/payments/bill-payment"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BillPaymentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing bill payment response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Receive Money
func (c *Client) ReceiveMoney(req ReceiveMoneyRequest) (*ReceiveMoneyResponse, error) {
	endpoint := "/payments/receive"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response ReceiveMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing receive money response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}
	
	// Set expiry date if not provided but expiry minutes is set in request
	if response.ExpiryDate.IsZero() && req.ExpiryMinutes > 0 {
		response.ExpiryDate = time.Now().UTC().Add(time.Duration(req.ExpiryMinutes) * time.Minute)
	}

	return &response, nil
}

// Transaction Query
func (c *Client) QueryTransaction(req TransactionQueryRequest) (*TransactionQueryResponse, error) {
	endpoint := "/transactions/status"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response TransactionQueryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing transaction query response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Helper function to validate amount is positive
func validateAmount(amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return fmt.Errorf("amount must be greater than zero")
	}
	return nil
}

// Airtime Purchase
func (c *Client) PurchaseAirtime(req AirtimePurchaseRequest) (*AirtimePurchaseResponse, error) {
	endpoint := "/payments/airtime"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response AirtimePurchaseResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing airtime purchase response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Bulk Payments
func (c *Client) ProcessBulkPayment(req BulkPaymentRequest) (*BulkPaymentResponse, error) {
	endpoint := "/payments/bulk"
	
	// Validate each payment item amount
	for _, item := range req.Items {
		if err := validateAmount(item.Amount); err != nil {
			return nil, fmt.Errorf("invalid amount for payment to %s: %w", item.DestinationAccount, err)
		}
	}
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BulkPaymentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing bulk payment response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) GetBulkPaymentStatus(req BulkPaymentStatusRequest) (*BulkPaymentStatusResponse, error) {
	endpoint := "/payments/bulk/status"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BulkPaymentStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing bulk payment status response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Standing Orders/Recurring Payments
func (c *Client) CreateStandingOrder(req StandingOrderRequest) (*StandingOrderResponse, error) {
	endpoint := "/payments/standing-orders"
	
	// Validate amount
	if err := validateAmount(req.Amount); err != nil {
		return nil, err
	}
	
	// Validate dates
	if req.StartDate.Before(time.Now()) {
		return nil, fmt.Errorf("start date must be in the future")
	}
	
	if !req.EndDate.IsZero() && req.EndDate.Before(req.StartDate) {
		return nil, fmt.Errorf("end date must be after start date")
	}
	
	// Validate frequency
	switch req.Frequency {
	case FrequencyDaily, FrequencyWeekly, FrequencyMonthly, FrequencyYearly:
		// Valid frequency
	default:
		return nil, fmt.Errorf("invalid frequency: %s", req.Frequency)
	}
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response StandingOrderResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing standing order response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) GetStandingOrderStatus(req StandingOrderStatusRequest) (*StandingOrderStatusResponse, error) {
	endpoint := "/payments/standing-orders/status"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response StandingOrderStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing standing order status response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) CancelStandingOrder(req StandingOrderCancelRequest) (*StandingOrderCancelResponse, error) {
	endpoint := "/payments/standing-orders/cancel"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response StandingOrderCancelResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing standing order cancel response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) ListStandingOrders(req StandingOrderListRequest) (*StandingOrderListResponse, error) {
	endpoint := "/payments/standing-orders/list"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response StandingOrderListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing standing order list response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Beneficiary Management
func (c *Client) CreateBeneficiary(req BeneficiaryCreateRequest) (*BeneficiaryCreateResponse, error) {
	endpoint := "/beneficiaries"
	
	// Validate beneficiary type
	switch req.Type {
	case BeneficiaryTypeBank:
		if req.AccountNumber == "" || req.BankCode == "" {
			return nil, fmt.Errorf("bank beneficiary requires account number and bank code")
		}
	case BeneficiaryTypeMobile:
		if req.MobileNumber == "" {
			return nil, fmt.Errorf("mobile beneficiary requires mobile number")
		}
	case BeneficiaryTypeBiller:
		if req.BillerCode == "" || req.CustomerReference == "" {
			return nil, fmt.Errorf("biller beneficiary requires biller code and customer reference")
		}
	default:
		return nil, fmt.Errorf("invalid beneficiary type: %s", req.Type)
	}
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BeneficiaryCreateResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing beneficiary create response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) ListBeneficiaries(req BeneficiaryListRequest) (*BeneficiaryListResponse, error) {
	endpoint := "/beneficiaries/list"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BeneficiaryListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing beneficiary list response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) GetBeneficiary(req BeneficiaryGetRequest) (*BeneficiaryGetResponse, error) {
	endpoint := "/beneficiaries/get"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BeneficiaryGetResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing beneficiary get response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) UpdateBeneficiary(req BeneficiaryUpdateRequest) (*BeneficiaryUpdateResponse, error) {
	endpoint := "/beneficiaries/update"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BeneficiaryUpdateResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing beneficiary update response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) DeleteBeneficiary(req BeneficiaryDeleteRequest) (*BeneficiaryDeleteResponse, error) {
	endpoint := "/beneficiaries/delete"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response BeneficiaryDeleteResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing beneficiary delete response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Foreign Exchange
func (c *Client) GetForexRate(req ForexRateRequest) (*ForexRateResponse, error) {
	endpoint := "/forex/rates"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response ForexRateResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing forex rate response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) ProcessForexTransfer(req ForexTransferRequest) (*ForexTransferResponse, error) {
	endpoint := "/forex/transfer"
	
	// Validate amount
	if err := validateAmount(req.SourceAmount); err != nil {
		return nil, err
	}
	
	// Validate currencies
	if req.SourceCurrency == req.DestinationCurrency {
		return nil, fmt.Errorf("source and destination currencies must be different")
	}
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response ForexTransferResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing forex transfer response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) GetForexTransferStatus(req ForexTransferStatusRequest) (*ForexTransferStatusResponse, error) {
	endpoint := "/forex/status"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response ForexTransferStatusResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing forex transfer status response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

// Authentication Methods
func (c *Client) RequestOTP(req OTPRequest) (*OTPResponse, error) {
	endpoint := "/auth/otp/request"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response OTPResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing OTP request response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}
	
	// Set expiry time if not provided
	if response.ExpiryTime.IsZero() {
		// Default expiry of 10 minutes if not specified
		response.ExpiryTime = time.Now().UTC().Add(10 * time.Minute)
	}

	return &response, nil
}

func (c *Client) VerifyOTP(req OTPVerifyRequest) (*OTPVerifyResponse, error) {
	endpoint := "/auth/otp/verify"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response OTPVerifyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing OTP verification response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) AuthenticateTransaction(req TransactionAuthRequest) (*TransactionAuthResponse, error) {
	endpoint := "/auth/transaction"
	
	// Validate authentication method
	switch req.AuthMethod {
	case AuthMethodOTP, AuthMethodBiometric, AuthMethod2FA:
		// Valid authentication method
	default:
		return nil, fmt.Errorf("invalid authentication method: %s", req.AuthMethod)
	}
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response TransactionAuthResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing transaction authentication response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}

func (c *Client) RegisterDevice(req DeviceRegistrationRequest) (*DeviceRegistrationResponse, error) {
	endpoint := "/auth/device/register"
	
	respBody, err := c.SendRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}

	var response DeviceRegistrationResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing device registration response: %w", err)
	}
	
	// Set timestamp to current time if not provided
	if response.Timestamp.IsZero() {
		response.Timestamp = time.Now().UTC()
	}

	return &response, nil
}
