package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type Environment string

const (
	SANDBOX    Environment = "sandbox"
	PRODUCTION Environment = "production"

	tokenURL              = "/collection/token/"
	disbursementTokenURL  = "/disbursement/token/"
	remittanceTokenURL    = "/remittance/token/"
	
	collectionTokenKey   = "collection_token"
	disbursementTokenKey = "disbursement_token"
	remittanceTokenKey   = "remittance_token"
)

type Client struct {
	apiUser         string
	apiKey          string
	subscriptionKey string
	environment     Environment
	baseURL         string
	httpClient      *http.Client
	cache           *cache.Cache
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type APIUserRequest struct {
	ProviderCallbackHost string `json:"providerCallbackHost"`
}

type APIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

func New(apiUser, apiKey, subscriptionKey string, environment Environment) (*Client, error) {
	if apiUser == "" || apiKey == "" || subscriptionKey == "" {
		return nil, fmt.Errorf("apiUser, apiKey, and subscriptionKey are required")
	}

	baseURL := "https://sandbox.momodeveloper.mtn.com"
	if environment == PRODUCTION {
		baseURL = "https://proxy.momoapi.mtn.com"
	}

	c := cache.New(1*time.Hour, 10*time.Minute)

	return &Client{
		apiUser:         apiUser,
		apiKey:          apiKey,
		subscriptionKey: subscriptionKey,
		environment:     environment,
		baseURL:         baseURL,
		httpClient:      &http.Client{Timeout: 30 * time.Second},
		cache:           c,
	}, nil
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) GetCollectionToken() (string, error) {
	return c.getToken(tokenURL, collectionTokenKey)
}

func (c *Client) GetDisbursementToken() (string, error) {
	return c.getToken(disbursementTokenURL, disbursementTokenKey)
}

func (c *Client) GetRemittanceToken() (string, error) {
	return c.getToken(remittanceTokenURL, remittanceTokenKey)
}

func (c *Client) getToken(endpoint, cacheKey string) (string, error) {
	if token, found := c.cache.Get(cacheKey); found {
		return token.(string), nil
	}

	url := c.baseURL + endpoint
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+c.getBasicAuth())
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token request failed with status: %s", resp.Status)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	expiresIn := tokenResp.ExpiresIn
	if expiresIn == 0 {
		expiresIn = 3600 
	}
	c.cache.Set(cacheKey, tokenResp.AccessToken, time.Duration(expiresIn)*time.Second)

	return tokenResp.AccessToken, nil
}

func (c *Client) getBasicAuth() string {
	auth := c.apiUser + ":" + c.apiKey
	return base64Encode(auth)
}

func (c *Client) makeRequest(method, url, product string, payload interface{}, headers map[string]string) ([]byte, error) {
	var token string
	var err error

	switch product {
	case "collection":
		token, err = c.GetCollectionToken()
	case "disbursement":
		token, err = c.GetDisbursementToken()
	case "remittance":
		token, err = c.GetRemittanceToken()
	default:
		return nil, fmt.Errorf("invalid product: %s", product)
	}

	if err != nil {
		return nil, err
	}

	var reqBody []byte
	if payload != nil {
		reqBody, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request payload: %w", err)
		}
	}

	req, err := http.NewRequest(method, c.baseURL+url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-Target-Environment", string(c.environment))
	req.Header.Set("Ocp-Apim-Subscription-Key", c.subscriptionKey)
	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	var respBody bytes.Buffer
	_, err = respBody.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, respBody.String())
	}

	return respBody.Bytes(), nil
}

func base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
