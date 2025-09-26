package stripe

import (
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/invoice"
	"github.com/stripe/stripe-go/v82/invoiceitem"
)

type InvoiceParams struct {
	CustomerID    string            `json:"customer"`
	SubscriptionID string           `json:"subscription,omitempty"`
	Description    string            `json:"description,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	DueDate        int64             `json:"due_date,omitempty"`
	AutoAdvance    bool              `json:"auto_advance,omitempty"`
	CollectionMethod string          `json:"collection_method,omitempty"`
	Lines []*InvoiceLineItem `json:"lines,omitempty"`
}

type InvoiceLineItem struct {
	Description string  `json:"description"`
	Amount      int64   `json:"amount"`
	Quantity    int64   `json:"quantity,omitempty"`
	TaxRates    []*string `json:"tax_rates,omitempty"`
}

func (c *Client) CreateInvoice(params InvoiceParams) (*stripe.Invoice, error) {
	invoiceParams := &stripe.InvoiceParams{
		Customer:         stripe.String(params.CustomerID),
		Description:      stripe.String(params.Description),
		Metadata:         params.Metadata,
		DueDate:          stripe.Int64(params.DueDate),
		AutoAdvance:      stripe.Bool(params.AutoAdvance),
		CollectionMethod: stripe.String(params.CollectionMethod),
	}

	if params.SubscriptionID != "" {
		invoiceParams.Subscription = stripe.String(params.SubscriptionID)
	}

	for _, line := range params.Lines {
		itemParams := &stripe.InvoiceItemParams{
			Customer:    stripe.String(params.CustomerID),
			Description: stripe.String(line.Description),
			Amount:      stripe.Int64(line.Amount),
			Quantity:    stripe.Int64(line.Quantity),
		}
		if len(line.TaxRates) > 0 {
			itemParams.TaxRates = line.TaxRates
		}
		_, err := invoiceitem.New(itemParams)
		if err != nil {
			return nil, err
		}
	}

	return invoice.New(invoiceParams)
}

func (c *Client) GetInvoice(invoiceID string) (*stripe.Invoice, error) {
	return invoice.Get(invoiceID, nil)
}

func (c *Client) UpdateInvoice(invoiceID string, params InvoiceParams) (*stripe.Invoice, error) {
	updateParams := &stripe.InvoiceParams{
		Description:      stripe.String(params.Description),
		Metadata:         params.Metadata,
		DueDate:          stripe.Int64(params.DueDate),
		AutoAdvance:      stripe.Bool(params.AutoAdvance),
		CollectionMethod: stripe.String(params.CollectionMethod),
	}

	return invoice.Update(invoiceID, updateParams)
}

func (c *Client) FinalizeInvoice(invoiceID string) (*stripe.Invoice, error) {
	return invoice.FinalizeInvoice(invoiceID, nil)
}

func (c *Client) PayInvoice(invoiceID string) (*stripe.Invoice, error) {
	return invoice.Pay(invoiceID, nil)
}

func (c *Client) VoidInvoice(invoiceID string) (*stripe.Invoice, error) {
	return invoice.VoidInvoice(invoiceID, nil)
}
func (c *Client) ListInvoices(params *stripe.InvoiceListParams) *invoice.Iter {
	return invoice.List(params)
}

func (c *Client) SendInvoice(invoiceID string) (*stripe.Invoice, error) {
	return invoice.SendInvoice(invoiceID, nil)
}
func (c *Client) GetInvoiceLines(invoiceID string) *invoice.LineItemIter {
	params := &stripe.InvoiceListLinesParams{
		Invoice: stripe.String(invoiceID),
	}
	return invoice.ListLines(params)
}
