package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	SandboxBaseURL    = "https://api-sandbox.fnb.co.za/integration-channel"
	ProductionBaseURL = "https://api.fnb.co.za/integration-channel"
)

type Client struct {
	clientID     string
	clientSecret string
	apiKey       string
	baseURL      string
	httpClient   *http.Client
	accessToken  string
	tokenExpiry  time.Time
	tokenMu      sync.RWMutex
	h2hConfig    *H2HConfig
}

type H2HConfig struct {
	CertPath     string
	KeyPath      string
	FTPHost      string
	FTPUser      string
	FTPPassword  string
	FTPDirectory string
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	Environment  string
	BaseURL      string
	H2HConfig    *H2HConfig
}

func NewClient(config *ClientConfig) *Client {
	baseURL := config.BaseURL
	if baseURL == "" {
		if config.Environment == "production" {
			baseURL = ProductionBaseURL
		} else {
			baseURL = SandboxBaseURL
		}
	}

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	if config.H2HConfig != nil && config.H2HConfig.CertPath != "" {
		cert, err := tls.LoadX509KeyPair(config.H2HConfig.CertPath, config.H2HConfig.KeyPath)
		if err == nil {
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
	}

	return &Client{
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
		apiKey:       config.APIKey,
		baseURL:      baseURL,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig:     tlsConfig,
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		h2hConfig: config.H2HConfig,
	}
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope,omitempty"`
}

type ErrorResponse struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Message          string `json:"message"`
	Code             string `json:"code"`
	Status           int    `json:"status"`
}

func (e *ErrorResponse) Error() string {
	if e.ErrorDescription != "" {
		return fmt.Sprintf("FNB API error: %s - %s", e.ErrorCode, e.ErrorDescription)
	}
	if e.Message != "" {
		return fmt.Sprintf("FNB API error [%s]: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("FNB API error: %s", e.ErrorCode)
}

func (c *Client) Authenticate(ctx context.Context) error {
	c.tokenMu.RLock()
	if c.isTokenValid() {
		c.tokenMu.RUnlock()
		return nil
	}
	c.tokenMu.RUnlock()

	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	if c.isTokenValid() {
		return nil
	}

	payload := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     c.clientID,
		"client_secret": c.clientSecret,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal auth payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/oauth/token", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute auth request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("authentication failed with status %d", resp.StatusCode)
		}
		return &errResp
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return fmt.Errorf("failed to decode auth response: %w", err)
	}

	c.accessToken = authResp.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(authResp.ExpiresIn-60) * time.Second) // Refresh 60s early

	return nil
}

func (c *Client) isTokenValid() bool {
	return c.accessToken != "" && time.Now().Before(c.tokenExpiry)
}

func (c *Client) DoRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	if err := c.Authenticate(ctx); err != nil {
		return err
	}

	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(data)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
		}
		errResp.Status = resp.StatusCode
		return &errResp
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) setHeaders(req *http.Request) {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.httpClient
}

func (c *Client) GetBaseURL() string {
	return c.baseURL
}
