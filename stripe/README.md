# Stripe Integration

This package provides a Go client for integrating with Stripe's payment processing services. It wraps the official [stripe-go](https://github.com/stripe/stripe-go) SDK with a more ergonomic interface.

## Features

- Payment Processing
  - Create and manage payment intents
  - Handle customer payments
  - Process refunds
  - Support for 3D Secure authentication

- Customer Management
  - Create and update customers
  - Manage payment methods
  - Handle customer metadata

- Subscription & Billing
  - Create and manage subscriptions
  - Handle invoices and billing
  - Support recurring payments

- Webhook Support
  - Verify webhook signatures
  - Process Stripe events

## Installation

```bash
go get github.com/stripe/stripe-go/v82
```

## Usage

### Initialize Client

```go
import "payment-rails/stripe"

// Create client with default configuration
client := stripe.NewClientWithKey("your_api_key")

// Or with custom configuration
client := stripe.NewClient(stripe.Config{
    APIKey:          "your_api_key",
    Environment:     stripe.Sandbox, // or stripe.Production
    TelemetryEnabled: true,
})
```

### Process a Payment

```go
// Create a payment intent
paymentIntent, err := client.CreatePaymentIntent(stripe.PaymentIntentParams{
    Amount:             2000, // Amount in cents
    Currency:           "usd",
    PaymentMethodTypes: []string{"card"},
    Description:        "Example payment",
})

// Confirm the payment intent
confirmedIntent, err := client.ConfirmPaymentIntent(paymentIntent.ID, &stripe.PaymentIntentConfirmParams{
    PaymentMethod: "pm_card_visa",
})
```

### Handle Subscriptions

```go
// Create a subscription
subscription, err := client.CreateSubscription(stripe.SubscriptionParams{
    CustomerID: "cus_123",
    Items: []*stripe.SubscriptionItem{
        {
            Price:    "price_H5ggYwtDq4fbrJ",
            Quantity: 1,
        },
    },
})
```

### Process Refunds

```go
// Create a refund
refund, err := client.CreateRefund(stripe.RefundParams{
    PaymentIntentID: "pi_123",
    Amount:          1000, // Amount in cents
    Reason:          "requested_by_customer",
})
```

### Create and Manage Customers

```go
// Create a customer
customer, err := client.CreateCustomer(stripe.CustomerParams{
    Email:       "customer@example.com",
    Description: "Example customer",
    PaymentMethod: "pm_card_visa",
})

// Attach a payment method
paymentMethod, err := client.AttachPaymentMethod(customer.ID, "pm_card_visa")
```

## Error Handling

The package uses Stripe's error types for consistent error handling:

```go
if err != nil {
    if stripeErr, ok := err.(*stripe.Error); ok {
        switch stripeErr.Code {
        case stripe.ErrorCodeCardDeclined:
            // Handle declined card
        case stripe.ErrorCodeExpiredCard:
            // Handle expired card
        default:
            // Handle other errors
        }
    }
}
```

## Examples

See the [examples directory](examples/) for complete usage examples:

- `basic_usage.go` - Basic Stripe operations using the official SDK directly
- `webhook_server.go` - Webhook handling server with event processing

### Running Examples

```bash
# Basic usage example
go run examples/basic_usage.go

# Webhook server example
go run examples/webhook_server.go
```

## Testing

The package is compatible with stripe-mock for testing. See the [stripe-mock documentation](https://github.com/stripe/stripe-mock) for setup instructions.

## Best Practices

1. Always use the sandbox environment for testing
2. Store API keys securely using environment variables
3. Implement proper error handling
4. Use webhooks for asynchronous events
5. Follow PCI compliance guidelines when handling card data

## License

This package is part of the payment-rails project.
