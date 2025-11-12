package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	SandboxBaseURL    = "https://api-sandbox.standardbank.co.za"
	ProductionBaseURL = "https://api.standardbank.co.za"

	DefaultTimeout     = 30 * time.Second
	TokenRefreshBuffer = 60 * time.Second // Refresh token 60s before expiry
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
	logger       Logger
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	Environment  string
	BaseURL      string
	Timeout      time.Duration
	Logger       Logger
}

type Logger interface {
	Log(level, message string, fields map[string]interface{})
}

type DefaultLogger struct{}

func (l *DefaultLogger) Log(level, message string, fields map[string]interface{}) {
	log.Printf("[%s] %s: %+v", level, message, fields)
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

	timeout := config.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	logger := config.Logger
	if logger == nil {
		logger = &DefaultLogger{}
	}

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	return &Client{
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
		apiKey:       config.APIKey,
		baseURL:      baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig:     tlsConfig,
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		logger: logger,
	}
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope,omitempty"`
}

type ErrorResponse struct {
	Error            string            `json:"error,omitempty"`
	ErrorDescription string            `json:"error_description,omitempty"`
	Message          string            `json:"message,omitempty"`
	Code             string            `json:"code,omitempty"`
	Status           int               `json:"status,omitempty"`
	Details          map[string]string `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e.ErrorDescription != "" {
		return fmt.Sprintf("Standard Bank API error: %s - %s", e.Error, e.ErrorDescription)
	}
	if e.Message != "" {
		return fmt.Sprintf("Standard Bank API error [%s]: %s", e.Code, e.Message)
	}
	if e.Error != "" {
		return fmt.Sprintf("Standard Bank API error: %s", e.Error)
	}
	return fmt.Sprintf("Standard Bank API error: status %d", e.Status)
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

	// Double-check after acquiring write lock
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

	c.logger.Log("DEBUG", "Authenticating with Standard Bank API", map[string]interface{}{
		"url": req.URL.String(),
	})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Log("ERROR", "Authentication request failed", map[string]interface{}{
			"error": err.Error(),
		})
		return fmt.Errorf("failed to execute auth request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read auth response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err == nil {
			errResp.Status = resp.StatusCode
			c.logger.Log("ERROR", "Authentication failed", map[string]interface{}{
				"status": resp.StatusCode,
				"error":  errResp.Error(),
			})
			return &errResp
		}
		c.logger.Log("ERROR", "Authentication failed", map[string]interface{}{
			"status": resp.StatusCode,
			"body":   string(respBody),
		})
		return fmt.Errorf("authentication failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	var authResp AuthResponse
	if err := json.Unmarshal(respBody, &authResp); err != nil {
		return fmt.Errorf("failed to decode auth response: %w", err)
	}

	c.accessToken = authResp.AccessToken
	expirySeconds := time.Duration(authResp.ExpiresIn) * time.Second
	c.tokenExpiry = time.Now().Add(expirySeconds - TokenRefreshBuffer)

	c.logger.Log("INFO", "Successfully authenticated", map[string]interface{}{
		"expires_in": authResp.ExpiresIn,
	})

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

	c.logger.Log("DEBUG", "Making API request", map[string]interface{}{
		"method": method,
		"url":    url,
	})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Log("ERROR", "API request failed", map[string]interface{}{
			"method": method,
			"url":    url,
			"error":  err.Error(),
		})
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err == nil {
			errResp.Status = resp.StatusCode
			c.logger.Log("ERROR", "API request failed", map[string]interface{}{
				"status": resp.StatusCode,
				"error":  errResp.Error(),
			})
			return &errResp
		}
		c.logger.Log("ERROR", "API request failed", map[string]interface{}{
			"status": resp.StatusCode,
			"body":   string(respBody),
		})
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	c.logger.Log("DEBUG", "API request successful", map[string]interface{}{
		"method": method,
		"url":    url,
		"status": resp.StatusCode,
	})

	return nil
}

func (c *Client) setHeaders(req *http.Request) {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.apiKey != "" {
		req.Header.Set("X-API-Key", c.apiKey)
	}
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.httpClient
}

func (c *Client) GetBaseURL() string {
	return c.baseURL
}
