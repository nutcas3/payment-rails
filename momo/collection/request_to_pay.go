package collection

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	requestToPayPath = "/collection/v1_0/requesttopay"
)

// RequestToPay requests payment from customer.
//
// Successful request returns nil response body and nil error, unless handleStatusPolling param is set.
// Callback should be handled by the caller.
//
// See [RequestToPay] docs for more information.
//
// [RequestToPay]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=RequesttoPay
func (c Collection) RequestToPay(
	ctx context.Context,
	refID uuid.UUID,
	callbackURL string,
	handleStatusPolling bool,
	body types.RequestToPayInput,
) (*types.RequestToPayStatus, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	headers, err := c.getHeaders(ctx, map[string]string{
		callbackHeader: callbackURL,
		refHeader:      refID.String(),
	})
	if err != nil {
		return nil, err
	}

	err = c.backend.Call(
		ctx,
		http.MethodPost,
		requestToPayPath,
		headers,
		nil,
		body,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if handleStatusPolling {
		resp, err := c.pollStatus(ctx, refID)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}

	return nil, nil
}

// RequestToPayTransactionStatus gets the status of a request to pay transaction.
//
// See [RequestToPayTransactionStatus] docs for more information.
//
// [RequestToPayTransactionStatus]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=RequesttoPayTransactionStatus
func (c Collection) RequestToPayTransactionStatus(ctx context.Context, refID uuid.UUID) (*types.RequestToPayStatus, error) {
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
		requestToPayPath,
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

// Helper for continously polling status of a request to pay transaction
func (c Collection) pollStatus(ctx context.Context, refID uuid.UUID) (*types.RequestToPayStatus, error) {
	var (
		resp    *types.RequestToPayStatus
		err     error
		attempt = 0
	)

	for {
		resp, err = c.RequestToPayTransactionStatus(ctx, refID)
		if err != nil {
			return nil, err
		}

		if resp.Status == "SUCCESSFUL" {
			break
		} else if resp.Status == "FAILED" {
			reason := "unknown reason"
			if resp.Reason.Message != "" {
				reason = resp.Reason.Message
			} else if resp.Reason.Code != "" {
				reason = resp.Reason.Code
			}
			return nil, fmt.Errorf("transaction failed: %s", reason)
		}

		attempt++
		if attempt > 6 {
			return nil, fmt.Errorf("timed out waiting for transaction status")
		}

		// Exponential backoff with cap at 30 seconds
		delay := time.Second * time.Duration(1<<uint(attempt))
		if delay > 30*time.Second {
			delay = 30 * time.Second
		}
		jitter := time.Duration(rand.Int63n(int64(delay / 4)))
		time.Sleep(delay + jitter)
	}

	return resp, nil
}

// RequestToPayDeliveryNotification sends additional notification to end user.
//
// See [RequestToPayDeliveryNotification] docs for more information.
//
// [RequestToPayDeliveryNotification]: https://momodeveloper.mtn.com/API-collections#api=collection&operation=RequesttoPayDeliveryNotification
func (c Collection) RequestToPayDeliveryNotification(
	ctx context.Context,
	refID uuid.UUID,
	message string,
	language string,
) (*types.DeliveryNotification, error) {
	if refID == uuid.Nil {
		return nil, types.ErrRefIDRequired
	}

	if message == "" {
		return nil, errors.New("message is required")
	}

	headers, err := c.getHeaders(ctx, map[string]string{
		"Language":            language,
		"notificationMessage": message,
	})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/deliverynotification", requestToPayPath, refID)

	var resp types.DeliveryNotification

	err = c.backend.Call(
		ctx,
		http.MethodPost,
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
