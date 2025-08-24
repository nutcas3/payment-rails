package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	receiveMoneyEndpoint      = "/transaction-api/v3.0/remittance/receive"
	receiveMoneyQueryEndpoint = "/transaction-api/v3.0/remittance/transaction"
)

func (c *Client) ReceiveMoney(req ReceiveMoneyRequest) (*ReceiveMoneyResponse, error) {
	if req.MerchantCode == "" || req.MerchantAccount == "" || req.CustomerName == "" || req.Amount == "" || req.CurrencyCode == "" || req.Reference == "" || req.Description == "" {
		return nil, fmt.Errorf("merchantCode, merchantAccount, customerName, amount, currencyCode, reference, and description are required")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	signatureData := req.MerchantCode + req.MerchantAccount + req.Amount + req.CurrencyCode

	respBody, err := c.SendRequest(http.MethodPost, receiveMoneyEndpoint, requestBody, signatureData)
	if err != nil {
		return nil, err
	}

	var response ReceiveMoneyResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing receive money response: %w", err)
	}

	return &response, nil
}

func (c *Client) QueryReceiveMoneyTransaction(req ReceiveMoneyQueryRequest) (*ReceiveMoneyQueryResponse, error) {
	if req.MerchantCode == "" || req.TransactionID == "" {
		return nil, fmt.Errorf("merchantCode and transactionId are required")
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	signatureData := req.MerchantCode + req.TransactionID

	respBody, err := c.SendRequest(http.MethodPost, receiveMoneyQueryEndpoint, requestBody, signatureData)
	if err != nil {
		return nil, err
	}

	var response ReceiveMoneyQueryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing receive money query response: %w", err)
	}

	return &response, nil
}
