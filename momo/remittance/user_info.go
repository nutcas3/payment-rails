package remittance

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common/types"
)

var ErrAccHolderMsisdnRequired = errors.New("accountHolderMSISDN is required")

const (
	userInfoPath          = "/remittance/v1_0/accountholder/msisdn/%s/basicuserinfo"
	userInfoPathClone     = "/remittance/clone-671b0/v1_0/accountholder/msisdn/%s/basicuserinfo"
	userInfoPathv3        = "/remittance/v1_0/accountholder/msisdn/999%s999/basicuserinfo"
	userConsentInfoPath   = "/remittance/oauth2/v1_0/userinfo"
	validateAccHolderPath = "/remittance/v1_0/accountholder/%s/%s/active"
)

// GetBasicUserInfo is used to return personal information of the account holder.
//
// See [GetBasicUserInfo] docs for more information.
//
// [GetBasicUserInfo]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=GetBasicUserinfo
func (r Remittance) GetBasicUserInfo(ctx context.Context, accHolderMsisdn string) (*types.BasicUserInfo, error) {
	if accHolderMsisdn == "" {
		return nil, ErrAccHolderMsisdnRequired
	}

	headers, err := r.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.BasicUserInfo

	url := fmt.Sprintf(userInfoPath, accHolderMsisdn)

	err = r.backend.Call(
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

// GetBasicUserInfov3 is used to return personal information of the account holder.
//
// See [GetBasicUserInfo-v3] docs for more information.
//
// [GetBasicUserInfo-v3]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=671b4b855d145f9b8f15e836
func (r Remittance) GetBasicUserInfov3(ctx context.Context, accHolderMsisdn string) (*types.BasicUserInfo, error) {
	if accHolderMsisdn == "" {
		return nil, ErrAccHolderMsisdnRequired
	}

	headers, err := r.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.BasicUserInfo

	url := fmt.Sprintf(userInfoPathv3, accHolderMsisdn)

	err = r.backend.Call(
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

// GetBasicUserInfoClone is used to return personal information of the account holder.
//
// See [GetBasicUserInfoClone] docs for more information.
//
// [GetBasicUserInfoClone]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=671b099705fd58bc55c9bbca
func (r Remittance) GetBasicUserInfoClone(ctx context.Context, accHolderMsisdn string) (*types.BasicUserInfo, error) {
	if accHolderMsisdn == "" {
		return nil, ErrAccHolderMsisdnRequired
	}

	headers, err := r.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.BasicUserInfo

	url := fmt.Sprintf(userInfoPathClone, accHolderMsisdn)

	err = r.backend.Call(
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
// [GetUserInfoWithConsent]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=GetUserInfoWithConsent
func (r Remittance) GetUserInfoWithConsent(ctx context.Context) (*types.UserConsentInfo, error) {
	token, err := r.getOauth2Token(ctx)
	if err != nil {
		return nil, err
	}

	headers, err := r.getHeaders(ctx, map[string]string{
		authHeader: "Bearer " + token,
	})
	if err != nil {
		return nil, err
	}

	var resp types.UserConsentInfo

	err = r.backend.Call(
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
// [ValidateAccountHolderStatus]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=ValidateAccountHolderStatus
func (d Remittance) ValidateAccountHolderStatus(ctx context.Context, accHolderID, accHolderIDType string) (bool, error) {
	if accHolderID == "" || accHolderIDType == "" {
		return false, errors.New("accountHolderId and accountHolderIdType are required")
	}

	headers, err := d.getHeaders(ctx, nil)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf(validateAccHolderPath, accHolderIDType, accHolderID)

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
