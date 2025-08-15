package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type Environment string

const (
	SANDBOX Environment = "sandbox"
	PRODUCTION Environment = "production"

	authURL                = "/auth/oauth2/token"
	ussdPushURL            = "/merchant/v1/payments/"
	transactionStatusURL   = "/standard/v1/payments/"
	refundURL              = "/standard/v1/payments/refund"
	disburseURL            = "/standard/v1/disbursements/"
	disbursementStatusURL  = "/standard/v1/disbursements/"
	accountBalanceURL      = "/standard/v1/accounts/balance"
	
	authTokenCacheKey = "airtel_auth_token"
)

type Service struct {
	clientID       string
	clientSecret   string
	publicKey      string
	environment    Environment
	country        string
	currency       string
	baseURL        string
	httpClient     *http.Client
	cache          *cache.Cache
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func New(clientID, clientSecret, publicKey string, environment Environment, country, currency string) (*Service, error) {
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("clientID and clientSecret are required")
	}

	baseURL := "https://openapiuat.airtel.africa"
	if environment == PRODUCTION {
		baseURL = "https://openapi.airtel.africa"
	}

	c := cache.New(1*time.Hour, 10*time.Minute)

	return &Service{
		clientID:      clientID,
		clientSecret:  clientSecret,
		publicKey:     publicKey,
		environment:   environment,
		country:       country,
		currency:      currency,
		baseURL:       baseURL,
		httpClient:    &http.Client{Timeout: 30 * time.Second},
		cache:         c,
	}, nil
}

func (s *Service) SetHttpClient(httpClient *http.Client) {
	s.httpClient = httpClient
}

func (s *Service) GetAuthToken() (string, error) {
	if token, found := s.cache.Get(authTokenCacheKey); found {
		return token.(string), nil
	}

	url := s.baseURL + authURL
	payload := map[string]string{
		"client_id":     s.clientID,
		"client_secret": s.clientSecret,
		"grant_type":    "client_credentials",
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal auth request payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute auth request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth request failed with status: %s", resp.Status)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", fmt.Errorf("failed to decode auth response: %w", err)
	}

	expiresIn := authResp.ExpiresIn
	if expiresIn == 0 {
		expiresIn = 3600 // Default to 1 hour if not provided
	}
	s.cache.Set(authTokenCacheKey, authResp.AccessToken, time.Duration(expiresIn)*time.Second)

	return authResp.AccessToken, nil
}

func (s *Service) makeRequest(method, url string, payload interface{}) ([]byte, error) {
	token, err := s.GetAuthToken()
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

	req, err := http.NewRequest(method, s.baseURL+url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Country", s.country)
	req.Header.Set("X-Currency", s.currency)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := readResponseBody(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errResp map[string]interface{}
		if err := json.Unmarshal(respBody, &errResp); err == nil {
			return nil, fmt.Errorf("request failed with status: %s, error: %v", resp.Status, errResp)
		}
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	return respBody, nil
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	var respBody bytes.Buffer
	_, err := respBody.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return respBody.Bytes(), nil
}
