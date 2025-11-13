package collection

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"
)

const (
	balancePath = "/collection/v1_0/account/balance"
)

// getHeaders is a helper for building request headers
//
// TODO: Look into a way of making this helper more DRY if possible. As of now it is repeated in disbursement and remittance
// packages because of the need to get auth tokens.
func (c Collection) getHeaders(ctx context.Context, vals map[string]string) (http.Header, error) {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	// Standard headers for all requests
	headers := http.Header{
		authHeader:    []string{"Bearer " + token},
		envHeader:     []string{c.environment},
		subHeader:     []string{c.subscriptionKey},
		contentHeader: []string{"application/json"},
	}

	// Variable headers for each request
	for k, v := range vals {
		headers[k] = []string{v}
	}

	return headers, nil
}

// GetAccountBalance is used to get balance of own account.
//
// See [GetAccountBalance] docs for more info.
//
// [GetAccountBalance]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetAccountBalance
func (c Collection) GetAccountBalance(ctx context.Context) (*types.Balance, error) {
	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.Balance

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		balancePath,
		headers,
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetAccountBalanceInSpecificCurrency is used to get balance of own account.
//
// See [GetAccountBalanceInSpecificCurrency] docs for more info.
//
// [GetAccountBalanceInSpecificCurrency]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetAccountBalance
func (c Collection) GetAccountBalanceInSpecificCurrency(ctx context.Context, currency types.Currency) (*types.Balance, error) {
	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.Balance

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		balancePath,
		headers,
		&common.Params{
			Path: []string{string(currency)},
		},
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
