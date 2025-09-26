package stripe

import (
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/customer"
	"github.com/stripe/stripe-go/v82/paymentmethod"
)
type CustomerParams struct {
	Email       string            `json:"email,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	PaymentMethod string `json:"payment_method,omitempty"`
	InvoiceSettings *stripe.CustomerInvoiceSettingsParams `json:"invoice_settings,omitempty"`
}

func (c *Client) CreateCustomer(params CustomerParams) (*stripe.Customer, error) {
	customerParams := &stripe.CustomerParams{
		Email:       stripe.String(params.Email),
		Name:        stripe.String(params.Name),
		Description: stripe.String(params.Description),
		Metadata:    params.Metadata,
	}

	if params.PaymentMethod != "" {
		customerParams.PaymentMethod = stripe.String(params.PaymentMethod)
	}

	if params.InvoiceSettings != nil {
		customerParams.InvoiceSettings = params.InvoiceSettings
	}

	return customer.New(customerParams)
}

func (c *Client) GetCustomer(customerID string) (*stripe.Customer, error) {
	return customer.Get(customerID, nil)
}

func (c *Client) UpdateCustomer(customerID string, params CustomerParams) (*stripe.Customer, error) {
	updateParams := &stripe.CustomerParams{
		Email:       stripe.String(params.Email),
		Name:        stripe.String(params.Name),
		Description: stripe.String(params.Description),
		Metadata:    params.Metadata,
	}

	if params.InvoiceSettings != nil {
		updateParams.InvoiceSettings = params.InvoiceSettings
	}

	return customer.Update(customerID, updateParams)
}

func (c *Client) DeleteCustomer(customerID string) (*stripe.Customer, error) {
	return customer.Del(customerID, nil)
}

func (c *Client) ListCustomers(params *stripe.CustomerListParams) *customer.Iter {
	return customer.List(params)
}
func (c *Client) AttachPaymentMethod(customerID, paymentMethodID string) (*stripe.PaymentMethod, error) {
	attachParams := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customerID),
	}

	return paymentmethod.Attach(paymentMethodID, attachParams)
}

func (c *Client) DetachPaymentMethod(paymentMethodID string) (*stripe.PaymentMethod, error) {
	return paymentmethod.Detach(paymentMethodID, nil)
}
func (c *Client) SetDefaultPaymentMethod(customerID, paymentMethodID string) (*stripe.Customer, error) {
	updateParams := &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(paymentMethodID),
		},
	}

	return customer.Update(customerID, updateParams)
}
