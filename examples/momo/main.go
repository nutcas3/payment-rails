package main

import (
	"log"
	"os"

	"github.com/nutcas3/payment-rails/momo"
)

func main() {
	targetEnv := os.Getenv("TARGET_ENVIRONMENT")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	CollectionSubKey := os.Getenv("COLLECTION_SUBSCRIPTION_KEY")
	DisbursementSubKey := os.Getenv("DISBURSEMENT_SUBSCRIPTION_KEY")
	RemittanceSubKey := os.Getenv("REMITTANCE_SUBSCRIPTION_KEY")

	cfg := momo.ClientConfig{
		Environment:                 targetEnv,
		APIKey:                      apiKey,
		APISecret:                   apiSecret,
		CollectionSubscriptionKey:   CollectionSubKey,
		DisbursementSubscriptionKey: DisbursementSubKey,
		RemittanceSubscriptionKey:   RemittanceSubKey,
	}

	_, err := momo.New(cfg)
	if err != nil {
		log.Fatalf("error starting client: %v", err)
	}

	// Check the /collection, /disbursement, /remittance subfolders for examples
	// of how to make respective API calls, or how to instantiate client for a
	// specific product only i.e Collectio, Disbursement or Remittance only.
}
