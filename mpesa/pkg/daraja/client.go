package daraja

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
	SANDBOX Environment = "sandbox"
	PRODUCTION Environment = "production"

	authURL                 = "/oauth/v1/generate?grant_type=client_credentials"
	stkPushURL              = "/mpesa/stkpush/v1/processrequest"
	stkPushQueryURL         = "/mpesa/stkpushquery/v1/query"
	c2bRegisterURL          = "/mpesa/c2b/v1/registerurl"
	c2bSimulateURL          = "/mpesa/c2b/v1/simulate"
	b2cURL                  = "/mpesa/b2c/v1/paymentrequest"
	b2bURL                  = "/mpesa/b2b/v1/paymentrequest"
	accountBalanceURL       = "/mpesa/accountbalance/v1/query"
	transactionStatusURL    = "/mpesa/transactionstatus/v1/query"
	reversalURL             = "/mpesa/reversal/v1/request"
	
	authTokenCacheKey = "auth_token"
)

type Service struct {
	apiKey          string
	consumerSecret  string
	passKey         string
	environment     Environment
	baseURL         string
	httpClient      *http.Client
	cache           *cache.Cache
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}
func New(apiKey, consumerSecret, passKey string, environment Environment) (*Service, error) {
	if apiKey == "" || consumerSecret == "" || passKey == "" {
		return nil, fmt.Errorf("apiKey, consumerSecret, and passKey are required")
	}

	baseURL := "https://sandbox.safaricom.co.ke"
	if environment == PRODUCTION {
		baseURL = "https://api.safaricom.co.ke"
	}

	c := cache.New(1*time.Hour, 10*time.Minute)

	return &Service{
		apiKey:         apiKey,
		consumerSecret: consumerSecret,
		passKey:        passKey,
		environment:    environment,
		baseURL:        baseURL,
		httpClient:     &http.Client{Timeout: 30 * time.Second},
		cache:          c,
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
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create auth request: %w", err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(s.apiKey + ":" + s.consumerSecret))
	req.Header.Set("Authorization", "Basic "+auth)
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

	expiresIn := 3600
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

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	respBody, err := readResponseBody(resp)
	if err != nil {
		return nil, err
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
