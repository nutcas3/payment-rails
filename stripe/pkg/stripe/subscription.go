package stripe

import (
	"fmt"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/subscription"
)

type SubscriptionParams struct {
	CustomerID           string              `json:"customer"`
	Items                []*SubscriptionItem `json:"items"`
	DefaultPaymentMethod string              `json:"default_payment_method,omitempty"`
	TrialPeriodDays      int64               `json:"trial_period_days,omitempty"`
	Metadata             map[string]string   `json:"metadata,omitempty"`
	CancelAtPeriodEnd    bool                `json:"cancel_at_period_end,omitempty"`
	ProrationBehavior    string              `json:"proration_behavior,omitempty"`
}

type SubscriptionItem struct {
	Price    string `json:"price"`
	Quantity int64  `json:"quantity,omitempty"`
}
func (c *Client) CreateSubscription(params SubscriptionParams) (*stripe.Subscription, error) {
	subParams := &stripe.SubscriptionParams{
		Customer:          stripe.String(params.CustomerID),
		Metadata:          params.Metadata,
		TrialPeriodDays:   stripe.Int64(params.TrialPeriodDays),
		CancelAtPeriodEnd: stripe.Bool(params.CancelAtPeriodEnd),
		ProrationBehavior: stripe.String(params.ProrationBehavior),
	}

	if params.DefaultPaymentMethod != "" {
		subParams.DefaultPaymentMethod = stripe.String(params.DefaultPaymentMethod)
	}

	for _, item := range params.Items {
		subParams.Items = append(subParams.Items, &stripe.SubscriptionItemsParams{
			Price:    stripe.String(item.Price),
			Quantity: stripe.Int64(item.Quantity),
		})
	}

	return subscription.New(subParams)
}

func (c *Client) GetSubscription(subscriptionID string) (*stripe.Subscription, error) {
	return subscription.Get(subscriptionID, nil)
}

func (c *Client) UpdateSubscription(subscriptionID string, params SubscriptionParams) (*stripe.Subscription, error) {
	updateParams := &stripe.SubscriptionParams{
		Customer:          stripe.String(params.CustomerID),
		Metadata:          params.Metadata,
		CancelAtPeriodEnd: stripe.Bool(params.CancelAtPeriodEnd),
		ProrationBehavior: stripe.String(params.ProrationBehavior),
	}

	if params.DefaultPaymentMethod != "" {
		updateParams.DefaultPaymentMethod = stripe.String(params.DefaultPaymentMethod)
	}

	for _, item := range params.Items {
		updateParams.Items = append(updateParams.Items, &stripe.SubscriptionItemsParams{
			Price:    stripe.String(item.Price),
			Quantity: stripe.Int64(item.Quantity),
		})
	}

	return subscription.Update(subscriptionID, updateParams)
}

func (c *Client) CancelSubscription(subscriptionID string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
	return subscription.Cancel(subscriptionID, params)
}

func (c *Client) ListSubscriptions(params *stripe.SubscriptionListParams) *subscription.Iter {
	return subscription.List(params)
}
func (c *Client) GetUpcomingInvoice(subscriptionID string) (*stripe.Invoice, error) {
	// Note: The Stripe Go SDK v82 doesn't have a direct "upcoming invoice" function
	// like other SDKs (e.g., Node.js has retrieveUpcoming). 
	// This is a known inconsistency across Stripe SDKs.
	// 
	// To get upcoming invoice information, you would typically:
	// 1. Use the Stripe API directly with HTTP calls
	// 2. Use subscription.Get() to get subscription details and calculate manually
	// 3. Wait for Stripe to add this functionality to the Go SDK
	//
	// For now, return an error indicating this limitation
	return nil, fmt.Errorf("GetUpcomingInvoice is not directly supported in Stripe Go SDK v82. Use direct API calls or subscription details instead")
}
