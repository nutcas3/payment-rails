package stripe

import (
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

type PaymentIntentParams struct {
	Amount             int64             `json:"amount"`
	Currency           string            `json:"currency"`
	CustomerID         string            `json:"customer,omitempty"`
	PaymentMethod      string            `json:"payment_method,omitempty"`
	PaymentMethodTypes []string          `json:"payment_method_types,omitempty"`
	Confirm            bool              `json:"confirm,omitempty"`
	OffSession         bool              `json:"off_session,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	Description        string            `json:"description,omitempty"`
	ReceiptEmail       string            `json:"receipt_email,omitempty"`
	SetupFutureUsage   string            `json:"setup_future_usage,omitempty"`

	ReturnURL string `json:"return_url,omitempty"`
}

func (c *Client) CreatePaymentIntent(params PaymentIntentParams) (*stripe.PaymentIntent, error) {
	piParams := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(params.Amount),
		Currency:           stripe.String(string(params.Currency)),
		Description:        stripe.String(params.Description),
		Metadata:           params.Metadata,
		ReceiptEmail:       stripe.String(params.ReceiptEmail),
		SetupFutureUsage:   stripe.String(params.SetupFutureUsage),
	}

	if params.CustomerID != "" {
		piParams.Customer = stripe.String(params.CustomerID)
	}

	if params.PaymentMethod != "" {
		piParams.PaymentMethod = stripe.String(params.PaymentMethod)
	}

	if len(params.PaymentMethodTypes) > 0 {
		piParams.PaymentMethodTypes = stripe.StringSlice(params.PaymentMethodTypes)
	}

	if params.Confirm {
		piParams.Confirm = stripe.Bool(true)
	}

	if params.OffSession {
		piParams.OffSession = stripe.Bool(true)
	}

	if params.ReturnURL != "" {
		piParams.ReturnURL = stripe.String(params.ReturnURL)
	}

	return paymentintent.New(piParams)
}


func (c *Client) GetPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	return paymentintent.Get(paymentIntentID, nil)
}

func (c *Client) UpdatePaymentIntent(paymentIntentID string, params PaymentIntentParams) (*stripe.PaymentIntent, error) {
	updateParams := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(params.Amount),
		Currency:    stripe.String(string(params.Currency)),
		Description: stripe.String(params.Description),
		Metadata:    params.Metadata,
	}

	if params.CustomerID != "" {
		updateParams.Customer = stripe.String(params.CustomerID)
	}

	if params.PaymentMethod != "" {
		updateParams.PaymentMethod = stripe.String(params.PaymentMethod)
	}

	return paymentintent.Update(paymentIntentID, updateParams)
}

func (c *Client) ConfirmPaymentIntent(paymentIntentID string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error) {
	return paymentintent.Confirm(paymentIntentID, params)
}

func (c *Client) CancelPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	return paymentintent.Cancel(paymentIntentID, nil)
}

func (c *Client) ListPaymentIntents(params *stripe.PaymentIntentListParams) *paymentintent.Iter {
	return paymentintent.List(params)
}

func (c *Client) CapturePaymentIntent(paymentIntentID string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error) {
	return paymentintent.Capture(paymentIntentID, params)
}
