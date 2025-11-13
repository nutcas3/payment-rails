package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	prodURL    = "https://momodeveloper.mtn.com"
	sandboxURL = "https://sandbox.momodeveloper.mtn.com"
)

// CacheStore defines the methods we use from the cache.
type CacheStore interface {
	Get(string) (any, bool)
	Set(string, any, time.Duration)
}

// NewCache creates and returns a new cache.
func NewCache() *cache.Cache {
	return cache.New(1*time.Hour, 3*time.Minute)
}

type Backend interface {
	Call(ctx context.Context, method, path string, headers http.Header, params *Params, body, result any) error
}

type BackendConfig struct {
	// Environment is the API environment being used i.e. sandbox or production.
	Environment string

	// HTTPClient is an HTTP client instance to use when making API requests.
	//
	// If left unset, it'll be set to a default HTTP client for the package.
	HTTPClient *http.Client
}

// Params represents path and query paramters
type Params struct {
	Path  []string
	Query map[string]string
}

// BackendImpl is an instance of a backend used to access a group of API methods i.e. Collection, Disbursement etc.
type BackendImpl struct {
	url        string
	HTTPClient *http.Client
}

// Call is the method for invoking API calls.
func (b *BackendImpl) Call(
	ctx context.Context,
	method, path string,
	headers http.Header,
	params *Params,
	body, result any,
) error {
	var (
		payload []byte
		err     error
	)

	if body != nil && !(reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil()) {
		payload, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	req, err := b.NewRequest(ctx, method, path, headers, params, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("momosdk: request failed with error: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("momosdk: request failed with status: %s", resp.Status)
	}

	// If the endpoint returns empty body skip decoding
	if resp.StatusCode == http.StatusNoContent || resp.ContentLength == 0 {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// NewRequest is used by Call to build a HTTP request. It handles encoding parameters and
// attaching headers.
func (b *BackendImpl) NewRequest(
	ctx context.Context,
	method, path string,
	headers http.Header,
	params *Params,
	body *bytes.Buffer,
) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	reqURL, err := url.Parse(b.url + path)
	if err != nil {
		return nil, fmt.Errorf("momosdk: error parsing URL: %w", err)
	}

	if params != nil {
		if params.Query != nil {
			queryParams := url.Values{}

			for k, v := range params.Query {
				queryParams.Add(k, v)
			}

			reqURL.RawQuery = queryParams.Encode()
		}

		if params.Path != nil {
			pathSegments := []string{strings.TrimSuffix(reqURL.Path, "/")}
			for _, v := range params.Path {
				pathSegments = append(pathSegments, url.PathEscape(v))
			}
			reqURL.Path = strings.Join(pathSegments, "/")
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("momosdk: error creating new request: %w", err)
	}

	for k, v := range headers {
		for _, item := range v {
			req.Header.Set(k, item)
		}
	}

	return req, nil
}

// NewBackend initializes a Backend using the config given.
func NewBackend(config *BackendConfig) (Backend, error) {
	cfg := *config
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{
			Timeout: 80 * time.Second,
		}
	}

	var baseURL string

	if cfg.Environment == "production" {
		baseURL = prodURL
	} else {
		baseURL = sandboxURL
	}

	return &BackendImpl{
		url:        baseURL,
		HTTPClient: cfg.HTTPClient,
	}, nil
}
