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

	qrResponse, err := client.GenerateQRCode(mpesa.QRCodeRequest{
		MerchantName: "TEST Business",               // Merchant Name
		RefNo:        "Invoice123",                  // Reference Number
		Amount:       1000,                          // Amount (in smallest currency unit)
		TrxCode:      mpesa.TrxCodeBuyGoods,         // Transaction Code (BG, WA, PB, SM, SB)
		CPI:          "254708374149",                // Customer Phone/Till/Paybill
		Size:         "300",                         // QR Code size in pixels
	})
	if err != nil {
		log.Fatalf("Failed to generate QR code: %v", err)
	}

	fmt.Printf("QR Code Response:\n")
	fmt.Printf("Response Code: %s\n", qrResponse.ResponseCode)
	fmt.Printf("Response Description: %s\n", qrResponse.ResponseDescription)
	fmt.Printf("QR Code (base64): %s...\n", qrResponse.QRCode[:50]) // Show just the beginning of the base64 string

	// Note: The QR Code is returned as a base64-encoded string in qrResponse.QRCode
	// You can convert this to an image and display it in your application
	fmt.Println("\nTo use this QR code in a web application, you can embed it like this:")
	fmt.Println("<img src=\"data:image/png;base64," + qrResponse.QRCode + "\" alt=\"M-Pesa QR Code\" />")
}
