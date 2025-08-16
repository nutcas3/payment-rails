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
	SANDBOX    Environment = "sandbox"
	PRODUCTION Environment = "production"

	accountInfoURL      = "/api/v1/account/info"
	accountBalanceURL   = "/api/v1/account/balance"
	accountStatementURL = "/api/v1/account/statement"
	accountTransferURL  = "/api/v1/account/transfer"
	
	forexRatesURL       = "/api/v1/forex/rates"
	forexExchangeURL    = "/api/v1/forex/exchange"
	
	voomaPayURL         = "/api/v1/vooma/pay"
	voomaStatusURL      = "/api/v1/vooma/status"
	pesalinkURL         = "/api/v1/pesalink/transfer"
	pesalinkStatusURL   = "/api/v1/pesalink/status"
	
	mobileMoneyURL      = "/api/v1/mobile/transfer"
	mobileMoneyStatusURL = "/api/v1/mobile/status"
	
	utilityPaymentURL   = "/api/v1/utility/pay"
	utilityStatusURL    = "/api/v1/utility/status"
	utilityProvidersURL = "/api/v1/utility/providers"
)

type Service struct {
	token        string
	environment  Environment
	baseURL      string
	httpClient   *http.Client
	cache        *cache.Cache
}

func New(token string, environment Environment) (*Service, error) {
	if token == "" {
		return nil, fmt.Errorf("token is required")
	}

	baseURL := "https://sandbox.buni.kcbgroup.com"
	if environment == PRODUCTION {
		baseURL = "https://buni.kcbgroup.com"
	}

	c := cache.New(1*time.Hour, 10*time.Minute)

	return &Service{
		token:       token,
		environment: environment,
		baseURL:     baseURL,
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		cache:       c,
	}, nil
}

func (s *Service) SetHttpClient(httpClient *http.Client) {
	s.httpClient = httpClient
}

func (s *Service) makeRequest(method, url string, payload interface{}) ([]byte, error) {
	var reqBody []byte
	var err error

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

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("Content-Type", "application/json")

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
