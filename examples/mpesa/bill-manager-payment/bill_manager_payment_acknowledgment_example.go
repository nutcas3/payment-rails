package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/nutcas3/payment-rails/mpesa"
	"time"
)

func main() {
	apiKey := os.Getenv("MPESA_API_KEY")
	consumerSecret := os.Getenv("MPESA_CONSUMER_SECRET")
	passKey := os.Getenv("MPESA_PASS_KEY")

	if apiKey == "" || consumerSecret == "" || passKey == "" {
		log.Fatal("Please set MPESA_API_KEY, MPESA_CONSUMER_SECRET, and MPESA_PASS_KEY environment variables")
	}

	client, err := mpesa.NewClient(
		apiKey,
		consumerSecret,
		passKey,
		mpesa.SANDBOX, // Use SANDBOX for testing, PRODUCTION for live environment
	)
	if err != nil {
		log.Fatalf("Failed to initialize Mpesa client: %v", err)
	}

	client.SetHttpClient(&http.Client{
		Timeout: 30 * time.Second,
	})

	// Send payment acknowledgment
	// This would typically be called after receiving a payment notification
	// at your callback URL
	response, err := client.SendPaymentAcknowledgment(mpesa.BillManagerAcknowledgmentRequest{
		PaymentDate:       "2023-09-15 14:30:45", // Date and time of payment
		PaidAmount:        "800",                 // Amount paid in KES
		AccountReference:  "1ASD678H",            // Account reference used in payment
		TransactionID:     "QXR12345678",         // M-PESA generated reference
		PhoneNumber:       "0722000000",          // Customer's phone number
		FullName:          "John Doe",            // Customer's full name
		InvoiceName:       "Monthly Subscription", // Name of the invoice
		ExternalReference: "#9932340",            // External reference of the invoice
	})
	if err != nil {
		log.Fatalf("Failed to send payment acknowledgment: %v", err)
	}

	fmt.Printf("Payment Acknowledgment Response:\n")
	fmt.Printf("Response Message: %s\n", response.ResMsg)
	fmt.Printf("Response Code: %s\n", response.ResCode)

	if response.ResCode == "200" {
		fmt.Println("\nPayment acknowledgment sent successfully!")
		fmt.Println("The payment has been reconciled in the Bill Manager system.")
		fmt.Println("The customer may receive a receipt based on your configuration.")
	} else {
		fmt.Println("\nFailed to send payment acknowledgment. Please check your request parameters.")
	}

	// Note: In a real application, you would implement a callback handler
	// to receive payment notifications from Safaricom, for example:
	/*
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Parse the incoming payment notification
		var paymentNotification mpesa.BillManagerPaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&paymentNotification); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Process the payment notification
		// ...

		// Send acknowledgment
		_, err := client.SendPaymentAcknowledgment(mpesa.BillManagerAcknowledgmentRequest{
			// Fill in the details from the payment notification
			// ...
		})
		if err != nil {
			log.Printf("Failed to send payment acknowledgment: %v", err)
		}

		// Respond to Safaricom
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"resmsg":  "success",
			"rescode": "200",
		})
	})
	*/
}
