package disbursement

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	refundPathv1 = "/disbursement/v1_0/refund"
	refundPathv2 = "/disbursement/v2_0/refund"
)

// RefundV1 is used to refund an amount from the owner's account to a payee account.
//
// See [RefundV1] docs for more info.
//
// [RefundV1]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=Refund-V2
func (d Disbursement) RefundV1(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RefundInput) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	return d.refund(ctx, refID.String(), callbackURL, refundPathv1, body)
}

// RefundV2 is used to refund an amount from the owner's account to a payee account.
//
// See [RefundV2] docs for more info.
//
// [RefundV2]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=Refund-V2
func (d Disbursement) RefundV2(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RefundInput) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	return d.refund(ctx, refID.String(), callbackURL, refundPathv2, body)
}

// refund is a helper for Refund requests
func (d Disbursement) refund(ctx context.Context, refID, callbackURL, path string, body types.RefundInput) error {
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

// GetRefundStatus is used to get the status of a refund.
//
// See [GetRefundStatus] docs for more info.
//
// [GetRefundStatus]: https://momodeveloper.mtn.com/API-collections#api=disbursement&operation=GetRefundStatus
func (d Disbursement) GetRefundStatus(ctx context.Context, refID uuid.UUID) (*types.DisbursementTransactionStatus, error) {
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
		refundPathv1,
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
