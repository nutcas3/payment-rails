package remittance

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/nutcas3/payment-rails/momo/common/types"
)

const (
	authTokenKey    = "remittanceAuthToken"
	oauth2TokenKey  = "remittanceOauth2Token"
	bcAuthKey       = "remittanceBcAuthToken"
	accessTokenPath = "/remittance/token/"
	oauth2Path      = "/remittance/oauth2/token/"
	bcAuthPath      = "/remittance/v1_0/bc-authorize"
)

// Helper for getting access token.
func (r Remittance) getAccessToken(ctx context.Context) (string, error) {
	if token, ok := r.cache.Get(authTokenKey); ok {
		return token.(string), nil
	}

	auth := base64.StdEncoding.EncodeToString([]byte(r.apiKey + ":" + r.apiSecret))

	headers := http.Header{
		contentHeader: []string{"application/json"},
		authHeader:    []string{"Basic " + auth},
	}

	var resp types.AccessTokenResp

	err := r.backend.Call(
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

	r.cache.Set(authTokenKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AccessToken, nil
}

// Helper for getting Oauth2 token.
func (r Remittance) getOauth2Token(ctx context.Context) (string, error) {
	if token, ok := r.cache.Get(oauth2TokenKey); ok {
		return token.(string), nil
	}

	auth := base64.StdEncoding.EncodeToString([]byte(r.apiKey + ":" + r.apiSecret))

	headers := http.Header{
		contentHeader: []string{"application/json"},
		authHeader:    []string{"Basic " + auth},
		envHeader:     []string{r.environment},
	}

	var resp types.Oauth2Resp

	err := r.backend.Call(
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

	r.cache.Set(oauth2TokenKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AccessToken, nil
}

// CreateAccessToken is used to create an access token.
//
// See [CreateAccessToken] docs for more info
//
// [CreateAccessToken]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=CreateAccessToken
func (r Remittance) CreateAccessToken(ctx context.Context) (string, error) {
	token, err := r.getAccessToken(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateOauth2Token is used to claim a consent by the account holder for the requested scopes
//
// See [CreateOauth2Token] docs for more info
//
// [CreateOauth2Token]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=CreateOauth2Token
func (r Remittance) CreateOauth2Token(ctx context.Context) (string, error) {
	token, err := r.getOauth2Token(ctx)
	if err != nil {
		return "", err
	}

	return token, nil
}

// BcAuthorize is used to claim a consent by the account holder for the requested scopes.
//
// See [BcAuthorize] docs for more info
//
// [BcAuthorize]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=bc-authorize
func (r Remittance) BcAuthorize(ctx context.Context, callbackURL string) (string, error) {
	if token, ok := r.cache.Get(bcAuthKey); ok {
		return token.(string), nil
	}

	token, err := r.getOauth2Token(ctx)
	if err != nil {
		return "", err
	}

	headers := http.Header{
		authHeader:     []string{"Bearer " + token},
		envHeader:      []string{r.environment},
		subHeader:      []string{r.subscriptionKey},
		contentHeader:  []string{"application/json"},
		callbackHeader: []string{callbackURL},
	}

	var resp types.BcAuthResp

	err = r.backend.Call(
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

	r.cache.Set(bcAuthKey, resp, time.Duration(resp.ExpiresIn)*time.Second)

	return resp.AuthRequestID, nil
}
