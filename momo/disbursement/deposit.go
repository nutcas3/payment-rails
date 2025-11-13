package disbursement

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	depositPathv1 = "/disbursement/v1_0/deposit"
	depositPathv2 = "/disbursement/v2_0/deposit"
)

// DepositV1 is used to deposit an amount from the owner account to a payee account.
//
// See [DepositV1] docs for more info.
//
// [DepositV1]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=Deposit-V1
func (d Disbursement) DepositV1(ctx context.Context, refID uuid.UUID, callbackURL string, body types.TransferInput) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	return d.deposit(ctx, refID.String(), callbackURL, depositPathv1, body)
}

// DepositV2 is used to deposit an amount from the owner account to a payee account.
//
// See [DepositV2] docs for more info.
//
// [DepositV2]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=Deposit-V2
func (d Disbursement) DepositV2(ctx context.Context, refID uuid.UUID, callbackURL string, body types.TransferInput) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	return d.deposit(ctx, refID.String(), callbackURL, depositPathv2, body)
}

// deposit is a helper for making deposits
func (d Disbursement) deposit(ctx context.Context, refID, callbackURL, path string, body types.TransferInput) error {
	headers, err := d.getHeaders(ctx, map[string]string{
		refHeader:      refID,
		callbackHeader: callbackURL,
	})
	if err != nil {
		return err
	}

	err = d.backend.Call(
		ctx,
		http.MethodPost,
		path,
		headers,
		nil,
		body,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetDepositStatus is used to get the status of a deposit.
//
// See [GetDepositStatus] docs for more info.
//
// [GetDepositStatus]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=GetDepositStatus
func (d Disbursement) GetDepositStatus(ctx context.Context, refID uuid.UUID) (*types.DisbursementTransactionStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := d.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.DisbursementTransactionStatus

	err = d.backend.Call(
		ctx,
		http.MethodGet,
		depositPathv1,
		headers,
		&common.Params{
			Path: []string{refID.String()},
		},
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
