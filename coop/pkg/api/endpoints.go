package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) AccountBalance(req AccountBalanceRequest) (*AccountBalanceResponse, error) {
	resp, err := c.makeRequest("POST", "/AccountBalance", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("account balance request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response AccountBalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode account balance response: %w", err)
	}

	return &response, nil
}

func (c *Client) AccountTransactions(req AccountTransactionsRequest) (*AccountTransactionsResponse, error) {
	resp, err := c.makeRequest("POST", "/AccountTransactions", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("account transactions request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response AccountTransactionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode account transactions response: %w", err)
	}

	return &response, nil
}

func (c *Client) ExchangeRate(req ExchangeRateRequest) (*ExchangeRateResponse, error) {
	resp, err := c.makeRequest("POST", "/ExchangeRate", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("exchange rate request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode exchange rate response: %w", err)
	}

	return &response, nil
}

func (c *Client) InternalFundsTransfer(req IFTRequest) (*IFTResponse, error) {
	resp, err := c.makeRequest("POST", "/IFTAccountToAccount", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("internal funds transfer request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response IFTResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode internal funds transfer response: %w", err)
	}

	return &response, nil
}

func (c *Client) PesaLinkSendToAccount(req PesaLinkRequest) (*PesaLinkResponse, error) {
	resp, err := c.makeRequest("POST", "/PesaLinkSendToAccount", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("pesalink transfer request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response PesaLinkResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode pesalink transfer response: %w", err)
	}

	return &response, nil
}

func (c *Client) TransactionStatus(req TransactionStatusRequest) (*TransactionStatusResponse, error) {
	resp, err := c.makeRequest("POST", "/TransactionStatus", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("transaction status request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response TransactionStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode transaction status response: %w", err)
	}

	return &response, nil
}
