package collection

import (
	"context"
	"net/http"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	paymentPath = "/collection/v2_0/payment"
)

// CreatePayments is used to perform payments.
//
// See [CreatePayments] docs for more info.
//
// [CreatePayments]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=CreatePayments
func (c Collection) CreatePayments(
	ctx context.Context,
	refID uuid.UUID,
	callbackURL string,
	body types.PaymentInput,
) error {
	if refID == uuid.Nil {
		return types.ErrRefIDRequired
	}

	headers, err := c.getHeaders(ctx, map[string]string{
		callbackHeader: callbackURL,
		refHeader:      refID.String(),
	})
	if err != nil {
		return err
	}

	err = c.backend.Call(
		ctx,
		http.MethodPost,
		paymentPath,
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

// GetPaymentStatus is used to get the status of a payment.
//
// See [GetPaymentStatus] docs for more info.
//
// [GetPaymentStatus]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=GetPaymentStatus
func (c Collection) GetPaymentStatus(ctx context.Context, refID uuid.UUID) (*types.PaymentStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := c.getHeaders(ctx, nil)
	if err != nil {
		return nil, err
	}

	var resp types.PaymentStatus

	err = c.backend.Call(
		ctx,
		http.MethodGet,
		paymentPath,
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
