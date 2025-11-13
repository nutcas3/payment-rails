package remittance

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"
)

const (
	balancePath = "/remittance/v1_0/account/balance"
)

// getHeaders is a helper for building request headers
//
// TODO: A way of making this helper more DRY. As is it is repeated in collection and disbursement
// packages because of the need to get auth tokens.
func (d Remittance) getHeaders(ctx context.Context, vals map[string]string) (http.Header, error) {
	token, err := d.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	// Standard headers for all requests
	headers := http.Header{
		authHeader:    []string{"Bearer " + token},
		envHeader:     []string{d.environment},
		subHeader:     []string{d.subscriptionKey},
		contentHeader: []string{"application/json"},
	}

	// Variable headers for each request
	for k, v := range vals {
		headers[k] = []string{v}
	}

	return headers, nil
}

// GetAccountBalance gets balance of own account.
//
// See [GetAccountBalance] docs for more information.
//
// [GetAccountBalance]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=GetAccountBalance
func (r Remittance) GetAccountBalance(ctx context.Context) (*types.Balance, error) {
	headers, err := r.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.Balance

	err = r.backend.Call(
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

// GetAccountBalanceInSpecificCurrency gets balance of own account.
//
// See [GetAccountBalanceInSpecificCurrency] docs for more information.
//
// [GetAccountBalanceInSpecificCurrency]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=GetAccountBalanceInSpecificCurrency
func (r Remittance) GetAccountBalanceInSpecificCurrency(ctx context.Context, currency types.Currency) (*types.Balance, error) {
	headers, err := r.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.Balance

	err = r.backend.Call(
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
