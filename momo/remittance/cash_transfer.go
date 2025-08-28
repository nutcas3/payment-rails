package remittance

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	cashTransferPath = "/remittance/v2_0/cashtransfer"
	transferPath     = "/remittance/v1_0/transfer"
)

// CashTransfer transfers money from the owner's account to a payee account.
//
// See [CashTransfer] docs for more info.
//
// [CashTransfer]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=CashTransfer
func (r Remittance) CashTransfer(
	ctx context.Context,
	refID uuid.UUID,
	callbackURL string,
	body types.CashTransferInput,
) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	headers, err := r.getHeaders(ctx, map[string]string{
		callbackHeader: callbackURL,
		refHeader:      refID.String(),
	})
	if err != nil {
		return err
	}

	err = r.backend.Call(
		ctx,
		http.MethodPost,
		cashTransferPath,
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

// GetCashTransferStatus gets the status of a cash transfer.
//
// See [GetCashTransferStatus] docs for more info.
//
// [GetCashTransferStatus]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=GetCashTransferStatus
func (r Remittance) GetCashTransferStatus(
	ctx context.Context,
	refID uuid.UUID,
) (*types.CashTransferStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := r.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.CashTransferStatus

	err = r.backend.Call(
		ctx,
		http.MethodGet,
		cashTransferPath,
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

// Transfer transfers an amount from own account to payee account.
//
// See [Transfer] docs for more info.
//
// [Transfer]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=Transfer
func (r Remittance) Transfer(
	ctx context.Context,
	refID uuid.UUID,
	callbackURL string,
	body types.TransferInput,
) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	headers, err := r.getHeaders(ctx, map[string]string{
		callbackHeader: callbackURL,
		refHeader:      refID.String(),
	})
	if err != nil {
		return err
	}

	err = r.backend.Call(
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

// GetTransferStatus get status of a transfer.
//
// See [GetTransferStatus] docs for more info.
//
// [GetTransferStatus]: https://momodeveloper.mtn.com/API-collections#api=remittance&operation=GetTransferStatus
func (r Remittance) GetTransferStatus(
	ctx context.Context,
	refID uuid.UUID,
) (*types.TransferStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := r.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.TransferStatus

	err = r.backend.Call(
		ctx,
		http.MethodGet,
		transferPath,
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
