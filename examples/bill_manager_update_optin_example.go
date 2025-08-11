package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-rails/mpesa"
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

	// Update Bill Manager opt-in details
	response, err := client.UpdateOptInDetails(mpesa.BillManagerUpdateOptInRequest{
		Shortcode:       "718003",                        // Your business shortcode
		Email:           "newemail@yourbusiness.com",     // Updated email address
		OfficialContact: "0710987654",                    // Updated contact phone number
		SendReminders:   1,                               // Enable reminders (1) or disable (0)
		Logo:            "updated_base64_encoded_image",  // Optional: Updated logo
		CallbackURL:     "https://newdomain.com/callback", // Updated callback URL
	})
	if err != nil {
		log.Fatalf("Failed to update opt-in details: %v", err)
	}

	fmt.Printf("Update Opt-In Details Response:\n")
	fmt.Printf("Response Message: %s\n", response.ResMsg)
	fmt.Printf("Response Code: %s\n", response.ResCode)

	if response.ResCode == "200" {
		fmt.Println("\nOpt-in details updated successfully!")
		fmt.Println("Your Bill Manager configuration has been updated with the new details.")
		fmt.Println("The changes include:")
		fmt.Println("- Updated email address")
		fmt.Println("- Updated contact phone number")
		fmt.Println("- Updated reminder settings")
		fmt.Println("- Updated callback URL")
		fmt.Println("- Updated logo (if provided)")
	} else {
		fmt.Println("\nFailed to update opt-in details. Please check your request parameters.")
	}
}
