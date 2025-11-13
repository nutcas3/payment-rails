package collection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common/types"
)

const (
	userInfoPath          = "/collection/v1_0/accountholder/%s/%s/basicuserinfo"
	userConsentInfoPath   = "/collection/oauth2/v1_0/userinfo"
	validateAccHolderPath = "/collection/v1_0/accountholder/%s/%s/active"
)

// GetBasicUserInfo is used to return personal information of the account holder.
//
// See [GetBasicUserInfo] docs for more information.
//
// [GetBasicUserInfo]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetBasicUserinfo
func (c Collection) GetBasicUserInfo(ctx context.Context, accHolderIDType, accHolderID string) (*types.BasicUserInfo, error) {
	if accHolderID == "" || accHolderIDType == "" {
		return nil, fmt.Errorf("accountHolderId and accountHolderIdType are required")
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.BasicUserInfo

	url := fmt.Sprintf(userInfoPath, accHolderIDType, accHolderID)

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		url,
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

// GetUserInfoWithConsent is used to return personal information of the account holder using their
// consent.
//
// See [GetUserInfoWithConsent] docs for more information.
//
// [GetUserInfoWithConsent]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetUserInfoWithConsent
func (c Collection) GetUserInfoWithConsent(ctx context.Context) (*types.UserConsentInfo, error) {
	token, err := c.getOauth2Token(ctx)
	if err != nil {
		return nil, err
	}

	headers, err := c.getHeaders(ctx, map[string]string{
		authHeader: "Bearer " + token,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UserConsentInfo

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		userConsentInfoPath,
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

// ValidateAccountHolderStatus is used to check if an account holder is registered and active in the system.
//
// See [ValidateAccountHolderStatus] docs for more information.
//
// [ValidateAccountHolderStatus]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=ValidateAccountHolderStatus
func (c Collection) ValidateAccountHolderStatus(ctx context.Context, accHolderID, accHolderIDType string) (bool, error) {
	if accHolderID == "" || accHolderIDType == "" {
		return false, fmt.Errorf("accountHolderId and accountHolderIdType are required")
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf(validateAccHolderPath, accHolderIDType, accHolderID)

	var status bool

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		url,
		headers,
		nil,
		nil,
		&status,
	)
	if err != nil {
		return false, err
	}

	return status, nil
}
