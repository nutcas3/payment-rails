package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURL = "https://api.ncbagroup.com" // This is a placeholder, replace with actual base URL
)

type Client struct {
	apiKey    string
	username  string
	password  string
	client    *http.Client
	apiToken  string
	tokenExp  time.Time
}

func NewClient(apiKey, username, password string) *Client {
	return &Client{
		apiKey:    apiKey,
		username:  username,
		password:  password,
		client:    &http.Client{Timeout: 30 * time.Second},
	}
}

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expiresIn"`
}

func (c *Client) Authenticate() error {
	if c.isTokenValid() {
		return nil
	}

	payload := map[string]string{
		"apiKey":   c.apiKey,
		"username": c.username,
		"password": c.password,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling auth payload: %v", err)
	}

	req, err := http.NewRequest("POST", BaseURL+"/auth", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error creating auth request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making auth request: %v", err)
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return fmt.Errorf("error decoding auth response: %v", err)
	}

	c.apiToken = authResp.Token
	c.tokenExp = time.Now().Add(time.Duration(authResp.ExpiresIn) * time.Second)
	return nil
}

func (c *Client) isTokenValid() bool {
	return c.apiToken != "" && time.Now().Before(c.tokenExp)
}

func (c *Client) setAuthHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.apiToken)
}
