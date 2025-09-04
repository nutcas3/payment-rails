package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	SandboxBaseURL = "https://sandbox.sasapay.app/api/v1"
	
	ProductionBaseURL = "https://api.sasapay.app/api/v1"
	
	tokenCacheKey = "sasapay_auth_token"
	
	tokenExpiryBuffer = 60
)

type Client struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
	HTTPClient   *http.Client
	TokenCache   *cache.Cache
	WebhookSecret string
}

func NewClient(clientID, clientSecret, environment string) (*Client, error) {
	baseURL := SandboxBaseURL
	if environment == "production" {
		baseURL = ProductionBaseURL
	}
	
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
		HTTPClient:   &http.Client{Timeout: 30 * time.Second},
		TokenCache:   NewCache(),
	}, nil
}

func NewCache() *cache.Cache {
	return cache.New(24*time.Hour, 1*time.Hour)
}

// SetWebhookSecret sets the secret used to validate webhook signatures
func (c *Client) SetWebhookSecret(secret string) {
	c.WebhookSecret = secret
}

// GetAuthToken retrieves an authentication token from the cache or from the API
func (c *Client) GetAuthToken() (string, error) {
	// Check if we have a cached token
	if token, found := c.TokenCache.Get(tokenCacheKey); found {
		return token.(string), nil
	}
	
	// No cached token, request a new one
	url := fmt.Sprintf("%s/auth/token", c.BaseURL)
	
	authReq := AuthTokenRequest{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
	}
	
	reqBody, err := json.Marshal(authReq)
	if err != nil {
		return "", fmt.Errorf("error marshalling auth request: %w", err)
	}
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error creating auth request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending auth request: %w", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading auth response: %w", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth request failed with status %d: %s", resp.StatusCode, string(body))
	}
	
	var authResp AuthTokenResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling auth response: %w", err)
	}
	
	// Cache the token
	// Set expiry to token expiry minus buffer to ensure we refresh before it expires
	expiryDuration := time.Duration(authResp.ExpiresIn-tokenExpiryBuffer) * time.Second
	c.TokenCache.Set(tokenCacheKey, authResp.AccessToken, expiryDuration)
	
	return authResp.AccessToken, nil
}

// SendRequest sends an HTTP request to the SasaPay API
func (c *Client) SendRequest(method, endpoint string, body interface{}) ([]byte, error) {
	// Get auth token
	token, err := c.GetAuthToken()
	if err != nil {
		return nil, fmt.Errorf("error getting auth token: %w", err)
	}
	
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	
	var reqBody []byte
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
	}
	
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiErr APIError
		if err := json.Unmarshal(respBody, &apiErr); err == nil {
			return nil, &apiErr
		}
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}
	
	return respBody, nil
}
