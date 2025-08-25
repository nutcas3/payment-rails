package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	SANDBOX    = "sandbox"
	PRODUCTION = "production"

	sandboxBaseURL    = "https://api-sandbox.absa.africa/v1"
	productionBaseURL = "https://api.absa.africa/v1"

	authEndpoint = "/oauth/token"

	tokenCacheKey = "absa_auth_token"
)

type Client struct {
	ClientID    string
	ClientSecret string
	APIKey      string
	Environment string
	BaseURL     string
	HTTPClient  *http.Client
	TokenCache  *cache.Cache
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient(clientID, clientSecret, apiKey, environment string) (*Client, error) {
	if clientID == "" || clientSecret == "" || apiKey == "" {
		return nil, errors.New("clientID, clientSecret, and apiKey are required")
	}

	baseURL := sandboxBaseURL
	if environment == PRODUCTION {
		baseURL = productionBaseURL
	}

	tokenCache := cache.New(50*time.Minute, 10*time.Minute)

	return &Client{
		ClientID:    clientID,
		ClientSecret: clientSecret,
		APIKey:      apiKey,
		Environment: environment,
		BaseURL:     baseURL,
		HTTPClient:  &http.Client{Timeout: 30 * time.Second},
		TokenCache:  tokenCache,
	}, nil
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.HTTPClient = httpClient
}

func (c *Client) GetAuthToken() (string, error) {
	if token, found := c.TokenCache.Get(tokenCacheKey); found {
		return token.(string), nil
	}

	reqBody := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling auth request: %w", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+authEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating auth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", c.APIKey)

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
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err == nil {
			return "", fmt.Errorf("auth error: %s (code: %d)", errResp.Message, errResp.Code)
		}
		return "", fmt.Errorf("auth error: %s", string(body))
	}

	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return "", fmt.Errorf("error parsing auth response: %w", err)
	}

	c.TokenCache.Set(tokenCacheKey, authResp.AccessToken, time.Duration(authResp.ExpiresIn)*time.Second)

	return authResp.AccessToken, nil
}

func (c *Client) SendRequest(method, endpoint string, body interface{}) ([]byte, error) {
	token, err := c.GetAuthToken()
	if err != nil {
		return nil, err
	}

	var jsonBody []byte
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request: %w", err)
		}
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Api-Key", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err == nil {
			return nil, fmt.Errorf("API error: %s (code: %d)", errResp.Message, errResp.Code)
		}
		return nil, fmt.Errorf("API error: %s", string(respBody))
	}

	return respBody, nil
}

func GenerateReference() string {
	timestamp := time.Now().Format("20060102150405")
	random := fmt.Sprintf("%06d", time.Now().Nanosecond()/1000)
	return timestamp + random[:6]
}
