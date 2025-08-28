package mpesa

import (
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

type BillManagerOptInRequest struct {
	Shortcode       string // Organization's shortcode (5-6 digits)
	Email           string // Official contact email address
	OfficialContact string // Official contact phone number
	SendReminders   string // "0" to disable, "1" to enable reminders
	Logo            string // Optional: Base64 encoded image for invoices/receipts
	CallbackURL     string // URL to receive payment notifications
}

type BillManagerOptInResponse struct {
	AppKey  string // App key received upon onboarding
	ResMsg  string // Response message
	ResCode string // Response code (200 for success)
}

type BillManagerInvoiceItem struct {
	ItemName string // Name of the invoice item
	Amount   string // Amount for this item
}

type BillManagerSingleInvoiceRequest struct {
	ExternalReference string                   // Unique invoice reference in your system
	BilledFullName    string                   // Name of the recipient
	BilledPhoneNumber string                   // Safaricom phone number to receive invoice
	BilledPeriod      string                   // Month and Year (e.g., "August 2021")
	InvoiceName       string                   // Descriptive name for what customer is being billed
	DueDate           string                   // Date customer is expected to pay (YYYY-MM-DD)
	AccountReference  string                   // Account number that uniquely identifies a customer
	Amount            string                   // Total invoice amount in KES
	InvoiceItems      []BillManagerInvoiceItem // Optional: Additional billable items
}

type BillManagerInvoiceResponse struct {
	StatusMessage string // Descriptive message
	ResMsg        string // Response message
	ResCode       string // Response code (200 for success)
}

type BillManagerCancelInvoiceRequest struct {
	ExternalReference string // External reference of invoice to cancel
}

type BillManagerCancelInvoiceResponse struct {
	StatusMessage string   // Descriptive message
	ResMsg        string   // Response message
	ResCode       string   // Response code (200 for success)
	Errors        []string // Any errors that occurred
}

type BillManagerPaymentRequest struct {
	TransactionID    string // M-PESA generated reference
	PaidAmount       string // Amount paid in KES
	MSISDN           string // Customer's phone number
	DateCreated      string // Date payment was recorded
	AccountReference string // Account reference used in payment
	ShortCode        string // Organization's shortcode
}

type BillManagerPaymentResponse struct {
	ResMsg  string // Response message
	ResCode string // Response code (200 for success)
}

type BillManagerAcknowledgmentRequest struct {
	PaymentDate       string // Date of payment
	PaidAmount        string // Amount paid in KES
	AccountReference  string // Account reference used in payment
	TransactionID     string // M-PESA generated reference
	PhoneNumber       string // Customer's phone number
	FullName          string // Customer's full name
	InvoiceName       string // Name of the invoice
	ExternalReference string // External reference of the invoice
}

type BillManagerUpdateOptInRequest struct {
	Shortcode       string // Organization's shortcode
	Email           string // Updated email address
	OfficialContact string // Updated contact phone number
	SendReminders   int    // 0 to disable, 1 to enable reminders
	Logo            string // Optional: Updated logo
	CallbackURL     string // Updated callback URL
}

func (c *Client) OptInBillManager(req BillManagerOptInRequest) (*BillManagerOptInResponse, error) {
	darajaReq := daraja.BillManagerOptInRequest{
		Shortcode:       req.Shortcode,
		Email:           req.Email,
		OfficialContact: req.OfficialContact,
		SendReminders:   req.SendReminders,
		Logo:            req.Logo,
		CallbackURL:     req.CallbackURL,
	}

	resp, err := c.Service.OptInBillManager(darajaReq)
	if err != nil {
		return nil, err
	}

	return &BillManagerOptInResponse{
		AppKey:  resp.AppKey,
		ResMsg:  resp.ResMsg,
		ResCode: resp.ResCode,
	}, nil
}

func (c *Client) CreateSingleInvoice(req BillManagerSingleInvoiceRequest) (*BillManagerInvoiceResponse, error) {
	invoiceItems := make([]daraja.BillManagerInvoiceItem, len(req.InvoiceItems))
	for i, item := range req.InvoiceItems {
		invoiceItems[i] = daraja.BillManagerInvoiceItem{
			ItemName: item.ItemName,
			Amount:   item.Amount,
		}
	}

	darajaReq := daraja.BillManagerSingleInvoiceRequest{
		ExternalReference: req.ExternalReference,
		BilledFullName:    req.BilledFullName,
		BilledPhoneNumber: req.BilledPhoneNumber,
		BilledPeriod:      req.BilledPeriod,
		InvoiceName:       req.InvoiceName,
		DueDate:           req.DueDate,
		AccountReference:  req.AccountReference,
		Amount:            req.Amount,
		InvoiceItems:      invoiceItems,
	}

	resp, err := c.Service.CreateSingleInvoice(darajaReq)
	if err != nil {
		return nil, err
	}

	return &BillManagerInvoiceResponse{
		StatusMessage: resp.StatusMessage,
		ResMsg:        resp.ResMsg,
		ResCode:       resp.ResCode,
	}, nil
}

func (c *Client) CreateBulkInvoices(requests []BillManagerSingleInvoiceRequest) (*BillManagerInvoiceResponse, error) {
	darajaReqs := make([]daraja.BillManagerSingleInvoiceRequest, len(requests))

	for i, req := range requests {
		invoiceItems := make([]daraja.BillManagerInvoiceItem, len(req.InvoiceItems))
		for j, item := range req.InvoiceItems {
			invoiceItems[j] = daraja.BillManagerInvoiceItem{
				ItemName: item.ItemName,
				Amount:   item.Amount,
			}
		}

		darajaReqs[i] = daraja.BillManagerSingleInvoiceRequest{
			ExternalReference: req.ExternalReference,
			BilledFullName:    req.BilledFullName,
			BilledPhoneNumber: req.BilledPhoneNumber,
			BilledPeriod:      req.BilledPeriod,
			InvoiceName:       req.InvoiceName,
			DueDate:           req.DueDate,
			AccountReference:  req.AccountReference,
			Amount:            req.Amount,
			InvoiceItems:      invoiceItems,
		}
	}

	resp, err := c.Service.CreateBulkInvoices(darajaReqs)
	if err != nil {
		return nil, err
	}

	return &BillManagerInvoiceResponse{
		StatusMessage: resp.StatusMessage,
		ResMsg:        resp.ResMsg,
		ResCode:       resp.ResCode,
	}, nil
}

func (c *Client) SendPaymentAcknowledgment(req BillManagerAcknowledgmentRequest) (*BillManagerPaymentResponse, error) {
	darajaReq := daraja.BillManagerAcknowledgmentRequest{
		PaymentDate:       req.PaymentDate,
		PaidAmount:        req.PaidAmount,
		AccountReference:  req.AccountReference,
		TransactionID:     req.TransactionID,
		PhoneNumber:       req.PhoneNumber,
		FullName:          req.FullName,
		InvoiceName:       req.InvoiceName,
		ExternalReference: req.ExternalReference,
	}

	resp, err := c.Service.SendPaymentAcknowledgment(darajaReq)
	if err != nil {
		return nil, err
	}

	return &BillManagerPaymentResponse{
		ResMsg:  resp.ResMsg,
		ResCode: resp.ResCode,
	}, nil
}

func (c *Client) CancelSingleInvoice(req BillManagerCancelInvoiceRequest) (*BillManagerCancelInvoiceResponse, error) {
	darajaReq := daraja.BillManagerCancelInvoiceRequest{
		ExternalReference: req.ExternalReference,
	}

	resp, err := c.Service.CancelSingleInvoice(darajaReq)
	if err != nil {
		return nil, err
	}

	return &BillManagerCancelInvoiceResponse{
		StatusMessage: resp.StatusMessage,
		ResMsg:        resp.ResMsg,
		ResCode:       resp.ResCode,
		Errors:        resp.Errors,
	}, nil
}

func (c *Client) CancelBulkInvoices(requests []BillManagerCancelInvoiceRequest) (*BillManagerCancelInvoiceResponse, error) {
	darajaReqs := make([]daraja.BillManagerCancelInvoiceRequest, len(requests))

	for i, req := range requests {
		darajaReqs[i] = daraja.BillManagerCancelInvoiceRequest{
			ExternalReference: req.ExternalReference,
		}
	}

	resp, err := c.Service.CancelBulkInvoices(darajaReqs)
	if err != nil {
		return nil, err
	}

	return &BillManagerCancelInvoiceResponse{
		StatusMessage: resp.StatusMessage,
		ResMsg:        resp.ResMsg,
		ResCode:       resp.ResCode,
		Errors:        resp.Errors,
	}, nil
}

func (c *Client) UpdateOptInDetails(req BillManagerUpdateOptInRequest) (*BillManagerPaymentResponse, error) {
	darajaReq := daraja.BillManagerUpdateOptInRequest{
		Shortcode:       req.Shortcode,
		Email:           req.Email,
		OfficialContact: req.OfficialContact,
		SendReminders:   req.SendReminders,
		Logo:            req.Logo,
		CallbackURL:     req.CallbackURL,
	}

	resp, err := c.Service.UpdateOptInDetails(darajaReq)
	if err != nil {
		return nil, err
	}

	return &BillManagerPaymentResponse{
		ResMsg:  resp.ResMsg,
		ResCode: resp.ResCode,
	}, nil
}
