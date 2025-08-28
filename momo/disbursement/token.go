package disbursement

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/nutcas3/payment-rails/momo/common/types"
)

const (
	authTokenKey    = "disbursementAuthToken"
	oauth2TokenKey  = "disbursementOauth2Token"
	bcAuthKey       = "disbursementBcAuthToken"
	accessTokenPath = "/disbursement/token/"
	oauth2Path      = "/disbursement/oauth2/token/"
	bcAuthPath      = "/disbursement/v1_0/bc-authorize"
)

// Helper for getting access token.
func (d Disbursement) getAccessToken(ctx context.Context) (string, error) {
	if token, ok := d.cache.Get(authTokenKey); ok {
		return token.(string), nil
	}

	auth := base64.StdEncoding.EncodeToString([]byte(d.apiKey + ":" + d.apiSecret))

	headers := http.Header{
		contentHeader: []string{"application/json"},
		authHeader:    []string{"Basic " + auth},
	}

	var resp types.AccessTokenResp

	err := d.backend.Call(
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

	d.cache.Set(authTokenKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AccessToken, nil
}

// Helper for getting Oauth2 token.
func (d Disbursement) getOauth2Token(ctx context.Context) (string, error) {
	if token, ok := d.cache.Get(oauth2TokenKey); ok {
		return token.(string), nil
	}

	auth := base64.StdEncoding.EncodeToString([]byte(d.apiKey + ":" + d.apiSecret))

	headers := http.Header{
		contentHeader: []string{"application/json"},
		authHeader:    []string{"Basic " + auth},
		envHeader:     []string{d.environment},
	}

	var resp types.Oauth2Resp

	err := d.backend.Call(
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

	d.cache.Set(oauth2TokenKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AccessToken, nil
}

// CreateAccessToken is used to create an access token.
//
// See [CreateAccessToken] docs for more info
//
// [CreateAccessToken]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=CreateAccessToken
func (d Disbursement) CreateAccessToken(ctx context.Context) (string, error) {
	token, err := d.getAccessToken(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateOauth2Token is used to claim a consent by the account holder for the requested scopes
//
// See [CreateOauth2Token] docs for more info
//
// [CreateOauth2Token]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=CreateOauth2Token
func (d Disbursement) CreateOauth2Token(ctx context.Context) (string, error) {
	token, err := d.getOauth2Token(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}

// BcAuthorize is used to claim a consent by the account holder for the requested scopes.
//
// See [BcAuthorize] docs for more info
//
// [BcAuthorize]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=bc-authorize
func (d Disbursement) BcAuthorize(ctx context.Context, callbackURL string) (string, error) {
	if token, ok := d.cache.Get(bcAuthKey); ok {
		return token.(string), nil
	}

	token, err := d.getOauth2Token(ctx)
	if err != nil {
		return "", err
	}

	headers := http.Header{
		authHeader:     []string{"Bearer " + token},
		envHeader:      []string{d.environment},
		subHeader:      []string{d.subscriptionKey},
		contentHeader:  []string{"application/json"},
		callbackHeader: []string{callbackURL},
	}

	var resp types.BcAuthResp

	err = d.backend.Call(
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

	d.cache.Set(bcAuthKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AuthRequestID, nil
}
