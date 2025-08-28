package disbursement

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common/types"
)

var ErrAccHolderIDAndTypeRequired = errors.New("accountHolderId and accountHolderIdType are required")

const (
	userInfoPath          = "/disbursement/v1_0/accountholder/%s/%s/basicuserinfo"
	userConsentInfoPath   = "/disbursement/oauth2/v1_0/userinfo"
	validateAccHolderPath = "/disbursement/v1_0/accountholder/%s/%s/active"
)

// GetBasicUserInfo is used to return personal information of the account holder.
//
// See [GetBasicUserInfo] docs for more information.
//
// [GetBasicUserInfo]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=GetBasicUserinfo
func (d Disbursement) GetBasicUserInfo(ctx context.Context, accHolderIDType, accHolderID string) (*types.BasicUserInfo, error) {
	if accHolderID == "" || accHolderIDType == "" {
		return nil, ErrAccHolderIDAndTypeRequired
	}

	headers, err := d.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.BasicUserInfo

	url := fmt.Sprintf(userInfoPath, accHolderIDType, accHolderID)

	err = d.backend.Call(
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
// [GetUserInfoWithConsent]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=GetUserInfoWithConsent
func (d Disbursement) GetUserInfoWithConsent(ctx context.Context) (*types.UserConsentInfo, error) {
	token, err := d.getOauth2Token(ctx)
	if err != nil {
		return nil, err
	}

	headers, err := d.getHeaders(ctx, map[string]string{
		authHeader: "Bearer " + token,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UserConsentInfo

	err = d.backend.Call(
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
// [ValidateAccountHolderStatus]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=ValidateAccountHolderStatus
func (d Disbursement) ValidateAccountHolderStatus(ctx context.Context, accHolderId, accHolderIdType string) (bool, error) {
	if accHolderId == "" || accHolderIdType == "" {
		return false, ErrAccHolderIDAndTypeRequired
	}

	headers, err := d.getHeaders(ctx, nil)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf(validateAccHolderPath, accHolderIdType, accHolderId)

	var status bool

	err = d.backend.Call(
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
