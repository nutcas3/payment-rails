package main

import (
	"fmt"
	"log"

	"payment-rails/stripe"
)

func main() {
	// Initialize the Stripe client with test API key
	client := stripe.NewClientWithKey("sk_test_your_api_key_here")

	fmt.Println("=== Payment Rails Stripe Example ===")

	// Example 1: Create a customer
	fmt.Println("\n1. Creating a Customer...")
	customer, err := client.CreateCustomer(stripe.CustomerParams{
		Email:       "customer@example.com",
		Name:        "John Doe",
		Description: "Example customer for testing",
		Metadata: map[string]string{
			"user_id": "12345",
		},
	})
	if err != nil {
		log.Printf("Error creating customer: %v", err)
		return
	}
	fmt.Printf("Created customer: %s\n", customer.ID)

	// Example 2: Create a payment intent
	fmt.Println("\n2. Creating a Payment Intent...")
	pi, err := client.CreatePaymentIntent(stripe.PaymentIntentParams{
		Amount:             2000, // $20.00 in cents
		Currency:           "usd",
		PaymentMethodTypes: []string{"card"},
		Description:        "Example payment",
		CustomerID:         customer.ID,
		Metadata: map[string]string{
			"order_id": "order_123",
		},
	})
	if err != nil {
		log.Printf("Error creating payment intent: %v", err)
		return
	}
	fmt.Printf("Created payment intent: %s (Status: %s)\n", pi.ID, pi.Status)
	fmt.Printf("Client secret: %s\n", pi.ClientSecret)

	// Example 3: Create a subscription
	fmt.Println("\n3. Creating a Subscription...")
	sub, err := client.CreateSubscription(stripe.SubscriptionParams{
		CustomerID: customer.ID,
		Items: []*stripe.SubscriptionItem{
			{
				Price:    "price_H5ggYwtDq4fbrJ", // Replace with actual price ID
				Quantity: 1,
			},
		},
		TrialPeriodDays: 7,
		Metadata: map[string]string{
			"plan": "premium",
		},
	})
	if err != nil {
		log.Printf("Error creating subscription: %v", err)
		return
	}
	fmt.Printf("Created subscription: %s (Status: %s)\n", sub.ID, sub.Status)

	// Example 4: Process a refund
	fmt.Println("\n4. Processing a Refund...")
	refund, err := client.CreateRefund(stripe.RefundParams{
		PaymentIntentID: pi.ID,
		Amount:          1000, // Partial refund of $10.00
		Reason:          "requested_by_customer",
		Metadata: map[string]string{
			"refund_reason": "customer_request",
		},
	})
	if err != nil {
		log.Printf("Error creating refund: %v", err)
		return
	}
	fmt.Printf("Created refund: %s (Status: %s, Amount: $%.2f)\n", 
		refund.ID, refund.Status, float64(refund.Amount)/100)

	// Example 5: Retrieve customer details
	fmt.Println("\n5. Retrieving Customer Details...")
	retrievedCustomer, err := client.GetCustomer(customer.ID)
	if err != nil {
		log.Printf("Error retrieving customer: %v", err)
		return
	}
	fmt.Printf("Retrieved customer: %s\n", retrievedCustomer.ID)
	fmt.Printf("  Name: %s\n", retrievedCustomer.Name)
	fmt.Printf("  Email: %s\n", retrievedCustomer.Email)
	fmt.Printf("  Created: %d\n", retrievedCustomer.Created)

	fmt.Println("\n=== Example Complete ===")
	fmt.Println("Note: Replace 'sk_test_your_api_key_here' with your actual Stripe test API key")
	fmt.Println("Note: Replace 'price_H5ggYwtDq4fbrJ' with an actual price ID from your Stripe dashboard")
}
