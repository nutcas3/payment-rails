package stripe

import (
	"fmt"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/refund"
)

type RefundParams struct {
	ChargeID    string            `json:"charge,omitempty"`
	PaymentIntentID string        `json:"payment_intent,omitempty"`
	Amount      int64             `json:"amount,omitempty"`
	Reason      string            `json:"reason,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

func (c *Client) CreateRefund(params RefundParams) (*stripe.Refund, error) {
	refundParams := &stripe.RefundParams{
		Metadata: params.Metadata,
		Reason:   stripe.String(params.Reason),
	}

	if params.Amount > 0 {
		refundParams.Amount = stripe.Int64(params.Amount)
	}

	if params.ChargeID != "" {
		refundParams.Charge = stripe.String(params.ChargeID)
	} else if params.PaymentIntentID != "" {
		refundParams.PaymentIntent = stripe.String(params.PaymentIntentID)
	} else {
		return nil, fmt.Errorf("either charge_id or payment_intent_id must be provided")
	}

	return refund.New(refundParams)
}

func (c *Client) GetRefund(refundID string) (*stripe.Refund, error) {
	return refund.Get(refundID, nil)
}

func (c *Client) UpdateRefund(refundID string, params RefundParams) (*stripe.Refund, error) {
	updateParams := &stripe.RefundParams{
		Metadata: params.Metadata,
	}

	return refund.Update(refundID, updateParams)
}

func (c *Client) ListRefunds(params *stripe.RefundListParams) *refund.Iter {
	return refund.List(params)
}
