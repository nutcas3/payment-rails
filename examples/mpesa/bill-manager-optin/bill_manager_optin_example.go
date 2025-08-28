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

	// Opt-in to Bill Manager
	response, err := client.OptInBillManager(mpesa.BillManagerOptInRequest{
		Shortcode:       "718003",                   // Your business shortcode
		Email:           "youremail@gmail.com",      // Official contact email
		OfficialContact: "0710123456",               // Official contact phone number
		SendReminders:   "1",                        // Enable reminders (1) or disable (0)
		Logo:            "base64_encoded_image",     // Optional: Base64 encoded image
		CallbackURL:     "https://example.com/callback", // URL to receive payment notifications
	})
	if err != nil {
		log.Fatalf("Failed to opt-in to Bill Manager: %v", err)
	}

	fmt.Printf("Bill Manager Opt-In Response:\n")
	fmt.Printf("App Key: %s\n", response.AppKey)
	fmt.Printf("Response Message: %s\n", response.ResMsg)
	fmt.Printf("Response Code: %s\n", response.ResCode)

	if response.ResCode == "200" {
		fmt.Println("\nSuccessfully opted in to Bill Manager!")
		fmt.Println("You can now use the Bill Manager features such as:")
		fmt.Println("- Creating single and bulk invoices")
		fmt.Println("- Receiving payments and sending acknowledgments")
		fmt.Println("- Canceling invoices")
		fmt.Println("- Updating your opt-in details")
	} else {
		fmt.Println("\nFailed to opt-in to Bill Manager. Please check your request parameters.")
	}
}
