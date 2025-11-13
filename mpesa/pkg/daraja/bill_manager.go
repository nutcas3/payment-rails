package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	billManagerOptInURL           = "/v1/billmanager-invoice/optin"
	billManagerSingleInvoicingURL = "/v1/billmanager-invoice/single-invoicing"
	billManagerBulkInvoicingURL   = "/v1/billmanager-invoice/bulk-invoicing"
	billManagerReconciliationURL  = "/v1/billmanager-invoice/reconciliation"
	billManagerCancelSingleURL    = "/v1/billmanager-invoice/cancel-single-invoice"
	billManagerCancelBulkURL      = "/v1/billmanager-invoice/cancel-bulk-invoices"
	billManagerUpdateOptInURL     = "/v1/billmanager-invoice/change-optin-details"
)

type BillManagerOptInRequest struct {
	Shortcode       string `json:"shortcode"`
	Email           string `json:"email"`
	OfficialContact string `json:"officialContact"`
	SendReminders   string `json:"sendReminders"` // "0" to disable, "1" to enable reminders
	Logo            string `json:"logo,omitempty"`
	CallbackURL     string `json:"callbackurl"`
}

type BillManagerOptInResponse struct {
	AppKey  string `json:"app_key"`
	ResMsg  string `json:"resmsg"`
	ResCode string `json:"rescode"`
}

type BillManagerInvoiceItem struct {
	ItemName string `json:"itemName"`
	Amount   string `json:"amount"`
}

type BillManagerSingleInvoiceRequest struct {
	ExternalReference string                   `json:"externalReference"`
	BilledFullName    string                   `json:"billedFullName"`
	BilledPhoneNumber string                   `json:"billedPhoneNumber"`
	BilledPeriod      string                   `json:"billedPeriod"`
	InvoiceName       string                   `json:"invoiceName"`
	DueDate           string                   `json:"dueDate"`
	AccountReference  string                   `json:"accountReference"`
	Amount            string                   `json:"amount"`
	InvoiceItems      []BillManagerInvoiceItem `json:"invoiceItems"`
}

type BillManagerInvoiceResponse struct {
	StatusMessage string `json:"Status_Message"`
	ResMsg        string `json:"resmsg"`
	ResCode       string `json:"rescode"`
}

type BillManagerCancelInvoiceRequest struct {
	ExternalReference string `json:"externalReference"`
}
type BillManagerCancelInvoiceResponse struct {
	StatusMessage string   `json:"Status_Message"`
	ResMsg        string   `json:"resmsg"`
	ResCode       string   `json:"rescode"`
	Errors        []string `json:"errors"`
}

type BillManagerPaymentRequest struct {
	TransactionID    string `json:"transactionId"`
	PaidAmount       string `json:"paidAmount"`
	MSISDN           string `json:"msisdn"`
	DateCreated      string `json:"dateCreated"`
	AccountReference string `json:"accountReference"`
	ShortCode        string `json:"shortCode"`
}

type BillManagerPaymentResponse struct {
	ResMsg  string `json:"resmsg"`
	ResCode string `json:"rescode"`
}

type BillManagerAcknowledgmentRequest struct {
	PaymentDate       string `json:"paymentDate"`
	PaidAmount        string `json:"paidAmount"`
	AccountReference  string `json:"accountReference"`
	TransactionID     string `json:"transactionId"`
	PhoneNumber       string `json:"phoneNumber"`
	FullName          string `json:"fullName"`
	InvoiceName       string `json:"invoiceName"`
	ExternalReference string `json:"externalReference"`
}

type BillManagerUpdateOptInRequest struct {
	Shortcode       string `json:"shortcode"`
	Email           string `json:"email"`
	OfficialContact string `json:"officialContact"`
	SendReminders   int    `json:"sendReminders"`
	Logo            string `json:"logo,omitempty"`
	CallbackURL     string `json:"callbackurl"`
}

func (s *Service) OptInBillManager(req BillManagerOptInRequest) (*BillManagerOptInResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, billManagerOptInURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to opt-in to Bill Manager: %w", err)
	}

	var response BillManagerOptInResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse Bill Manager opt-in response: %w", err)
	}

	return &response, nil
}

func (s *Service) CreateSingleInvoice(req BillManagerSingleInvoiceRequest) (*BillManagerInvoiceResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, billManagerSingleInvoicingURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create single invoice: %w", err)
	}

	var response BillManagerInvoiceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse single invoice response: %w", err)
	}

	return &response, nil
}

func (s *Service) CreateBulkInvoices(requests []BillManagerSingleInvoiceRequest) (*BillManagerInvoiceResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, billManagerBulkInvoicingURL, requests)
	if err != nil {
		return nil, fmt.Errorf("failed to create bulk invoices: %w", err)
	}

	var response BillManagerInvoiceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse bulk invoice response: %w", err)
	}

	return &response, nil
}

func (s *Service) SendPaymentAcknowledgment(req BillManagerAcknowledgmentRequest) (*BillManagerPaymentResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, billManagerReconciliationURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to send payment acknowledgment: %w", err)
	}

	var response BillManagerPaymentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse payment acknowledgment response: %w", err)
	}

	return &response, nil
}

func (s *Service) CancelSingleInvoice(req BillManagerCancelInvoiceRequest) (*BillManagerCancelInvoiceResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, billManagerCancelSingleURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel single invoice: %w", err)
	}

	var response BillManagerCancelInvoiceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse cancel single invoice response: %w", err)
	}

	return &response, nil
}

// CancelBulkInvoices cancels multiple invoices
func (s *Service) CancelBulkInvoices(requests []BillManagerCancelInvoiceRequest) (*BillManagerCancelInvoiceResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, billManagerCancelBulkURL, requests)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel bulk invoices: %w", err)
	}

	var response BillManagerCancelInvoiceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse cancel bulk invoices response: %w", err)
	}

	return &response, nil
}

// UpdateOptInDetails updates the Bill Manager opt-in details
func (s *Service) UpdateOptInDetails(req BillManagerUpdateOptInRequest) (*BillManagerPaymentResponse, error) {
	respBody, err := s.makeRequest(http.MethodPost, billManagerUpdateOptInURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update opt-in details: %w", err)
	}

	var response BillManagerPaymentResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse update opt-in details response: %w", err)
	}

	return &response, nil
}
