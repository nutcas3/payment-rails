package collection

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/nutcas3/payment-rails/momo/common/types"
)

const (
	authTokenKey    = "collectionAuthToken"
	oauth2TokenKey  = "collectionOauth2Token"
	bcAuthKey       = "collectionBcAuthToken"
	accessTokenPath = "/collection/token/"
	oauth2Path      = "/collection/oauth2/token/"
	bcAuthPath      = "/collection/v1_0/bc-authorize"
)

// Helper for getting access token.
func (c Collection) getAccessToken(ctx context.Context) (string, error) {
	if token, ok := c.cache.Get(authTokenKey); ok {
		return token.(string), nil
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.apiKey + ":" + c.apiSecret))

	headers := http.Header{
		contentHeader: []string{"application/json"},
		authHeader:    []string{"Basic " + auth},
	}

	var resp types.AccessTokenResp

	err := c.backend.Call(
		ctx,
		http.MethodPost,
		accessTokenPath,
		headers,
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return "", err
	}

	c.cache.Set(authTokenKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AccessToken, nil
}

// Helper for getting Oauth2 token.
func (c Collection) getOauth2Token(ctx context.Context) (string, error) {
	if token, ok := c.cache.Get(oauth2TokenKey); ok {
		return token.(string), nil
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.apiKey + ":" + c.apiSecret))

	headers := http.Header{
		contentHeader: []string{"application/json"},
		authHeader:    []string{"Basic " + auth},
		envHeader:     []string{c.environment},
	}

	var resp types.Oauth2Resp

	err := c.backend.Call(
		ctx,
		http.MethodPost,
		oauth2Path,
		headers,
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return "", err
	}

	c.cache.Set(oauth2TokenKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AccessToken, nil
}

// CreateAccessToken is used to create an access token.
//
// See [CreateAccessToken] docs for more information.
//
// [CreateAccessToken]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=CreateAccessToken
func (c Collection) CreateAccessToken(ctx context.Context) (string, error) {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateOauth2Token is used to claim a consent by the account holder for the requested scopes.
//
// See [CreateOauth2Token] docs for more information.
//
// [CreateOauth2Token]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=CreateOauth2Tokenbg
func (c Collection) CreateOauth2Token(ctx context.Context) (string, error) {
	token, err := c.getOauth2Token(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}

// BcAuthorize is used to claim a consent by the account holder for the requested scopes.
//
// See [bcAuthorize] docs for more information.
//
// [bcAuthorize]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=bc-authorize
func (c Collection) BcAuthorize(ctx context.Context, callbackURL string) (string, error) {
	if token, ok := c.cache.Get(bcAuthKey); ok {
		return token.(string), nil
	}

	token, err := c.getOauth2Token(ctx)
	if err != nil {
		return "", err
	}

	headers := http.Header{
		authHeader:     []string{"Bearer " + token},
		envHeader:      []string{c.environment},
		subHeader:      []string{c.subscriptionKey},
		contentHeader:  []string{"application/json"},
		callbackHeader: []string{callbackURL},
	}

	var resp types.BcAuthResp

	err = c.backend.Call(
		ctx,
		http.MethodPost,
		bcAuthPath,
		headers,
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return "", err
	}

	c.cache.Set(bcAuthKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AuthRequestID, nil
}
