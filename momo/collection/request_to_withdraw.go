package collection

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	withdrawPathv1 = "/collection/v1_0/requesttowithdraw"
	withdrawPathv2 = "/collection/v2_0/requesttowithdraw"
)

// RequestToWithdrawTransactionStatus is used to get the status of request to withdraw transaction.
//
// See [RequestToWithdrawTransactionStatus] docs for more info.
//
// [RequestToWithdrawTransactionStatus]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=RequestToWithdrawTransactionStatus
func (c Collection) RequestToWithdrawTransactionStatus(ctx context.Context, refID uuid.UUID) (*types.RequestToPayStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.RequestToPayStatus

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		withdrawPathv1,
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

// RequestToWithdrawV1 is used to request a withdrawal from a consumer.
//
// See [RequestToWithdrawV1] docs for more info.
//
// [RequestToWithdrawV1]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=RequestToWithdraw-V1
func (c Collection) RequestToWithdrawV1(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RequestToPayInput) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	return c.requestToWithdraw(ctx, refID.String(), callbackURL, withdrawPathv1, body)
}

// RequestToWithdrawV2 is used to request a withdrawal from a consumer.
//
// See [RequestToWithdrawV2] docs for more info.
//
// [RequestToWithdrawV2]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=RequestToWithdraw-V2
func (c Collection) RequestToWithdrawV2(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RequestToPayInput) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	return c.requestToWithdraw(ctx, refID.String(), callbackURL, withdrawPathv2, body)
}

// Helper for request to withdraw operation.
func (c Collection) requestToWithdraw(ctx context.Context, refID, callbackURL, path string, body types.RequestToPayInput) error {
	headers, err := c.getHeaders(ctx, map[string]string{
		refHeader:      refID,
		callbackHeader: callbackURL,
	})
	if err != nil {
		return err
	}

	err = c.backend.Call(
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
