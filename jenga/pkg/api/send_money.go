package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	remittanceEndpoint = "/transaction-api/v3.0/remittance"
	mobileWalletEndpoint = "/transaction-api/v3.0/remittance/mobile"
	rtgsEndpoint = "/transaction-api/v3.0/remittance/rtgs"
	swiftEndpoint = "/transaction-api/v3.0/remittance/swift"
)

func (c *Client) SendMoney(req SendMoneyRequest) (*SendMoneyResponse, error) {
	if req.Source.CountryCode == "" || req.Source.AccountNumber == "" || 
	   req.Destination.CountryCode == "" || req.Destination.AccountNumber == "" ||
	   req.Transfer.Amount == "" || req.Transfer.CurrencyCode == "" || req.Transfer.Reference == "" {
		return nil, fmt.Errorf("missing required fields in SendMoneyRequest")
	}

	endpoint := remittanceEndpoint
	switch req.Transfer.Type {
	case TransferTypeRTGS:
		endpoint = rtgsEndpoint
	case TransferTypeSWIFT:
		endpoint = swiftEndpoint
	}

	signatureData := req.Source.AccountNumber + req.Transfer.Amount + req.Transfer.CurrencyCode + req.Transfer.Reference

	respBody, err := c.SendRequest(http.MethodPost, endpoint, req, signatureData)
	if err != nil {
		return nil, err
	}

	var response SendMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing send money response: %w", err)
	}

	return &response, nil
}

func (c *Client) SendToMobileWallet(req MobileWalletRequest) (*MobileWalletResponse, error) {
	if req.Source.CountryCode == "" || req.Source.AccountNumber == "" || 
	   req.Destination.CountryCode == "" || req.Destination.MobileNumber == "" || req.Destination.WalletName == "" ||
	   req.Transfer.Amount == "" || req.Transfer.CurrencyCode == "" || req.Transfer.Reference == "" {
		return nil, fmt.Errorf("missing required fields in MobileWalletRequest")
	}

	signatureData := req.Source.AccountNumber + req.Transfer.Amount + req.Transfer.CurrencyCode + req.Transfer.Reference

	respBody, err := c.SendRequest(http.MethodPost, mobileWalletEndpoint, req, signatureData)
	if err != nil {
		return nil, err
	}

	var response MobileWalletResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing mobile wallet response: %w", err)
	}

	return &response, nil
}
