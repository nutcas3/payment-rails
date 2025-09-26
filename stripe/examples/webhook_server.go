package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	stripeSDK "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"
	"payment-rails/stripe"
)

var client *stripe.Client

// WebhookHandler handles Stripe webhook events
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// Replace with your webhook endpoint secret
	endpointSecret := "whsec_your_webhook_secret_here"
	
	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), endpointSecret)
	if err != nil {
		log.Printf("Error verifying webhook signature: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Handle the event
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripeSDK.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handlePaymentIntentSucceeded(&paymentIntent)
		
	case "payment_intent.payment_failed":
		var paymentIntent stripeSDK.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handlePaymentIntentFailed(&paymentIntent)
		
	case "customer.subscription.created":
		var subscription stripeSDK.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handleSubscriptionCreated(&subscription)
		
	default:
		log.Printf("Unhandled event type: %s", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}

func handlePaymentIntentSucceeded(pi *stripeSDK.PaymentIntent) {
	log.Printf("✅ Payment succeeded for PaymentIntent: %s", pi.ID)
	log.Printf("   Amount: $%.2f", float64(pi.Amount)/100)
	if pi.Customer != nil {
		// Get customer details using our SDK
		customer, err := client.GetCustomer(pi.Customer.ID)
		if err != nil {
			log.Printf("   Error fetching customer: %v", err)
		} else {
			log.Printf("   Customer: %s (%s)", customer.ID, customer.Email)
		}
	}
	
	// Add your business logic here
	// For example:
	// - Update order status in database
	// - Send confirmation email
	// - Fulfill the order
}

func handlePaymentIntentFailed(pi *stripeSDK.PaymentIntent) {
	log.Printf("❌ Payment failed for PaymentIntent: %s", pi.ID)
	log.Printf("   Amount: $%.2f", float64(pi.Amount)/100)
	if pi.Customer != nil {
		// Get customer details using our SDK
		customer, err := client.GetCustomer(pi.Customer.ID)
		if err != nil {
			log.Printf("   Error fetching customer: %v", err)
		} else {
			log.Printf("   Customer: %s (%s)", customer.ID, customer.Email)
		}
	}
	
	// Add your business logic here
	// For example:
	// - Update order status in database
	// - Send payment failure notification
	// - Retry payment or cancel order
}

func handleSubscriptionCreated(sub *stripeSDK.Subscription) {
	log.Printf("🎉 Subscription created: %s", sub.ID)
	
	// Get customer details using our SDK
	customer, err := client.GetCustomer(sub.Customer.ID)
	if err != nil {
		log.Printf("   Error fetching customer: %v", err)
	} else {
		log.Printf("   Customer: %s (%s)", customer.ID, customer.Email)
	}
	log.Printf("   Status: %s", sub.Status)
	
	// Add your business logic here
	// For example:
	// - Grant access to premium features
	// - Send welcome email
	// - Update user permissions
}

func main() {
	// Initialize Stripe client
	client = stripe.NewClientWithKey("sk_test_your_api_key_here")

	// Set up webhook endpoint
	http.HandleFunc("/webhook", WebhookHandler)
	
	fmt.Println("🚀 Payment Rails Stripe Webhook Server Starting...")
	fmt.Println("📡 Webhook endpoint: http://localhost:8080/webhook")
	fmt.Println("")
	fmt.Println("⚠️  Remember to:")
	fmt.Println("   1. Replace 'sk_test_your_api_key_here' with your actual Stripe API key")
	fmt.Println("   2. Replace 'whsec_your_webhook_secret_here' with your actual webhook secret")
	fmt.Println("   3. Configure your Stripe webhook endpoint to point to this server")
	fmt.Println("   4. Select the events you want to receive in your Stripe dashboard")
	fmt.Println("")
	fmt.Println("📋 Supported events:")
	fmt.Println("   - payment_intent.succeeded")
	fmt.Println("   - payment_intent.payment_failed") 
	fmt.Println("   - customer.subscription.created")
	fmt.Println("")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
