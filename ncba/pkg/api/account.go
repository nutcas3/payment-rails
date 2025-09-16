package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountDetails struct {
	AccountNumber string  `json:"accountNumber"`
	AccountName   string  `json:"accountName"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
	Status        string  `json:"status"`
}

type MiniStatement struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Balance     float64 `json:"balance"`
}

type AccountStatement struct {
	Transactions []Transaction `json:"transactions"`
	Period      struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"period"`
}

type Transaction struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Reference   string  `json:"reference"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Balance     float64 `json:"balance"`
}

func (c *Client) GetAccountDetails(countryCode, accountNo string) (*AccountDetails, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/accounts/%s/%s", BaseURL, countryCode, accountNo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(req)
	
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result AccountDetails
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}

func (c *Client) GetMiniStatement(countryCode, accountNo string) ([]MiniStatement, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/accounts/%s/%s/mini-statement", BaseURL, countryCode, accountNo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result []MiniStatement
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return result, nil
}

func (c *Client) GetAccountStatement(countryCode, accountNo, fromDate, toDate string) (*AccountStatement, error) {
	if err := c.Authenticate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/accounts/%s/%s/statement?from=%s&to=%s", BaseURL, countryCode, accountNo, fromDate, toDate)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	c.setAuthHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result AccountStatement
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}
