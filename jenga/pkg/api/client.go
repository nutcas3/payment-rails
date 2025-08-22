package api

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	SANDBOX = "sandbox"
	PRODUCTION = "production"

	sandboxBaseURL    = "https://uat.finserve.africa/v3-apis"
	productionBaseURL = "https://api.finserve.africa/v3-apis"

	authEndpoint = "/authentication/api/v3/authenticate/merchant"

	tokenCacheKey = "jenga_auth_token"
)

type Client struct {
	APIKey      string
	Username    string
	Password    string
	PrivateKey  string
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

func NewClient(apiKey, username, password, privateKey, environment string) (*Client, error) {
	if apiKey == "" || username == "" || password == "" || privateKey == "" {
		return nil, errors.New("apiKey, username, password, and privateKey are required")
	}

	baseURL := sandboxBaseURL
	if environment == PRODUCTION {
		baseURL = productionBaseURL
	}

	tokenCache := cache.New(50*time.Minute, 10*time.Minute)

	return &Client{
		APIKey:      apiKey,
		Username:    username,
		Password:    password,
		PrivateKey:  privateKey,
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
		"username": c.Username,
		"password": c.Password,
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

func (c *Client) GenerateSignature(data string) (string, error) {
	block, _ := pem.Decode([]byte(c.PrivateKey))
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	hash := sha256.Sum256([]byte(data))

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func (c *Client) SendRequest(method, endpoint string, body interface{}, signatureData string) ([]byte, error) {
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

	if signatureData != "" {
		signature, err := c.GenerateSignature(signatureData)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Signature", signature)
	}

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
