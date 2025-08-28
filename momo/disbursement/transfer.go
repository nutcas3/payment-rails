package disbursement

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	transferPath = "/disbursement/v1_0/transfer"
)

// Transfer is used to transfer amount from own account to payee account.
//
// See [Transfer] docs for more info.
//
// [Transfer]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=Transfer
func (d Disbursement) Transfer(
	ctx context.Context,
	refID uuid.UUID,
	callbackURL string,
	body types.TransferInput,
) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	headers, err := d.getHeaders(ctx, map[string]string{
		refHeader:      refID.String(),
		callbackHeader: callbackURL,
	})
	if err != nil {
		return err
	}

	err = d.backend.Call(
		ctx,
		http.MethodPost,
		transferPath,
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

// GetTransferStatus is used to get the status of a transfer.
//
// See [GetTransferStatus] docs for more info.
//
// [GetTransferStatus]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=GetTransferStatus
func (d Disbursement) GetTransferStatus(
	ctx context.Context,
	refID uuid.UUID,
) (*types.DisbursementTransactionStatus, error) {
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
		transferPath,
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
